package clients

import (
	"Savings/utils/cache"
	"Savings/utils/logger"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
	"time"
)

type ApiClient struct {
	client  *http.Client
	request *http.Request
	baseUrl string
	cache   cache.ICache[string, string]
}

type AuthResponse struct {
	Token string `json:"access_token"`
}

var clientCache cache.ICache[string, string]

func Init() {
	logger.Log.Debug("Init client")

	clientCache = cache.New[string, string]()
}

type ApiResponse struct {
	Result  int         `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func New(baseUrl string) *ApiClient {
	logger.Log.Debug("New client: ", "url", baseUrl)

	return &ApiClient{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseUrl: baseUrl,
		cache:   clientCache,
	}
}

func (api *ApiClient) getUrl(endpoint string) string {
	if strings.HasPrefix(endpoint, "http") {
		return endpoint
	}
	if !strings.HasPrefix(api.baseUrl, "http") {
		api.baseUrl = "https://" + api.baseUrl
	}
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}
	return api.baseUrl + endpoint
}

func (api *ApiClient) Send(data interface{}) error {
	//TODO: Can we encode the data for security purposes and decode when necessary? Same to response logging...
	logger.Log.Debug("API_REQ: ", "req", api.request)
	start := time.Now()
	response, err := api.client.Do(api.request)
	if err != nil {
		logger.Log.Error("Error sending request to API endpoint: ", err)
		return err
	}

	// Close the connection to reuse it
	defer response.Body.Close()
	logger.Log.Info("API_RES - raw: ", "res", response, "lat", time.Since(start))

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Log.Debug("Couldn't parse response body: ", err)
	}
	logger.Log.Info("API_RES - body: ", "body", body)

	//TODO: Perform error handling in a better way
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		var errRes ApiResponse
		if err = json.NewDecoder(response.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		// TODO: Add retries for Auth
		if response.StatusCode == 401 {
			panic("Failed to authenticate.")
		}

		return fmt.Errorf("unknown error, status code: %d", response.StatusCode)
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		logger.Log.Debug("Failed to unmarshal body: ", err)
	}

	return nil
}

func (api *ApiClient) setDefaultHeaders() {
	api.request.Header = http.Header{
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
	}
}

func (api *ApiClient) baseRequest(method string, endpoint string, body io.Reader) *ApiClient {
	endpoint = api.getUrl(endpoint)
	request, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		logger.Log.Error("error creating HTTP request: %v", err)
	}

	api.request = request
	api.setDefaultHeaders()

	return api
}

func (api *ApiClient) NewRequest(method string, endpoint string, body io.Reader) *ApiClient {
	if token := api.cache.GetString("token"); token != "" {
		api.baseRequest(method, endpoint, body).request.Header.Add("Authorization", "Bearer "+token)
	} else {
		api.ensureAuthenticated()

		token = api.cache.GetString("token")
		api.baseRequest(method, endpoint, body).request.Header.Add("Authorization", "Bearer "+token)
	}

	return api
}

func (api *ApiClient) ensureAuthenticated() {
	values := map[string]string{"email": "aa@a.a", "password": "12345678"}
	jsonData, err := json.Marshal(values)

	err = api.authenticate(jsonData)
	if err != nil {
		logger.Log.Error("error authenticating: %v", err)
	}
}

func (api *ApiClient) authenticate(data []byte) error {
	var response = new(AuthResponse)

	err := api.baseRequest(http.MethodPost, viper.GetString("SIDOOH_ACCOUNTS_API_URL")+"/users/signin", bytes.NewBuffer(data)).Send(response)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if api.cache != nil {
		api.cache.Set("token", response.Token, 14*time.Minute)
	}

	return nil
}
