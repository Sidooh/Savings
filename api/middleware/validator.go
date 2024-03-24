package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator *validator.Validate

func validate(i interface{}) error {
	if err := Validator.Struct(i); err != nil {
		//var validationErrors []responses.ValidationError

		//for _, err := range err.(validator.ValidationErrors) {
		//
		//	//msg := fmt.Sprintf("%v is invalid/missing", err.Field())
		//
		//	tag := err.Tag()
		//	if err.Param() != "" {
		//		tag += " " + err.Param()
		//	}
		//
		//	//validationErrors = append(validationErrors, responses.ValidationError{
		//	//	Value:   err.Value(),
		//	//	Field:   err.Field(),
		//	//	Message: msg,
		//	//	Param:   tag,
		//	//})
		//}

		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func BindAndValidateRequest(context *fiber.Ctx, request interface{}) error {
	if err := context.BodyParser(request); err != nil {
		return err
	}

	if err := validate(request); err != nil {
		return err
	}

	return nil
}
