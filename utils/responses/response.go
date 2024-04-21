package responses

import (
	"Savings/ent"
	"Savings/utils"
	internal_errors "Savings/utils/errors"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

func ErrorResponse(err interface{}, c *fiber.Ctx) JsonResponse {
	r := JsonResponse{
		Status:  false,
		Message: "an error occurred",
		Errors:  nil,
	}
	c.Status(fiber.StatusInternalServerError)

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
		c.Status(fiber.StatusUnprocessableEntity)
	case error:
		var NFE *ent.NotFoundError
		var CE *ent.ValidationError

		switch {
		case errors.As(t, &NFE):
			r.Message = "not found"
			c.Status(fiber.StatusNotFound)
		case errors.As(t, &CE):
			c.Status(fiber.StatusUnprocessableEntity)
			r.Message = "validation failed for db update"
		case errors.Is(t, internal_errors.InsufficientBalance):
			c.Status(fiber.StatusPaymentRequired)
			r.Message = t.Error()
		case errors.Is(t, internal_errors.InvalidIdParameter):
			c.Status(fiber.StatusUnprocessableEntity)
			r.Message = t.Error()
		default:
			r.Message = t.Error()
		}
	default:
		r.Errors = t
	}

	return r
}
