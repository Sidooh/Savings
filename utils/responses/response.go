package responses

import (
	"Savings/ent"
	"Savings/utils"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

type JsonResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"` // application-level error messages, for debugging
	Data    interface{} `json:"data,omitempty"`   // data wrapper
	Meta    interface{} `json:"meta,omitempty"`   // meta wrapper e.g. pagination data
}

type ValidationError struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Param   string      `json:"param"`
	Value   interface{} `json:"value,omitempty"`
}

func (j JsonResponse) Error() string {
	v, _ := jsoniter.MarshalToString(j)
	return v
}

func SuccessResponse(data interface{}) JsonResponse {
	return JsonResponse{
		Status: true,
		Data:   data,
	}
}

func PaginatedResponse(data interface{}, paginator *utils.Paginator) JsonResponse {
	return JsonResponse{
		Status: true,
		Data:   data,
		Meta:   paginator.Meta(),
	}
}

func ErrorResponse(err interface{}) JsonResponse {
	r := JsonResponse{
		Status:  false,
		Message: "an error occurred",
		Errors:  nil,
	}

	switch t := err.(type) {
	case validator.ValidationErrors:
		m := make(map[string]string, len(t))
		for _, err := range t {
			switch err.Tag() {
			case "required":
				m[err.Field()] = "This value is required."
			case "email":
				m[err.Field()] = "This is not a valid email."
			default:
				s := fmt.Sprintf("This value failed validation on '%s", err.Tag())
				if err.Param() != "" {
					s += fmt.Sprintf(":%s'.", err.Param())
				} else {
					s += "'."
				}
				m[err.Field()] = s
			}
		}
		r.Errors = m
		r.Message = "invalid request"
	case error:
		var NFE *ent.NotFoundError
		var CE *ent.ValidationError
		switch {
		case errors.As(t, &NFE):
			r.Message = "not found"
		case errors.As(t, &CE):
			r.Message = "validation failed for db update"
		default:
			r.Message = t.Error()
		}
	default:
		r.Errors = t
	}

	return r
}
