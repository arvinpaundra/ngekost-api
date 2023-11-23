package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError map[string]string

func Validate(value any) ValidationError {
	validate := validator.New()

	// register tag json to be validated, instead validate struct property
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		// skip tag if want to be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(value)

	if err != nil {
		return formatErrorValidation(err)
	}

	return nil
}

func formatErrorValidation(validationError error) ValidationError {
	errFields := make(ValidationError)

	// make error message for each invalid field
	for _, err := range validationError.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errFields[err.Field()] = "this field is required"
		case "email":
			errFields[err.Field()] = "invalid email format"
		case "min":
			errFields[err.Field()] = fmt.Sprintf("min length %s characters", err.Param())
		case "max":
			errFields[err.Field()] = fmt.Sprintf("max length %s characters", err.Param())
		case "required_if":
			errFields[err.Field()] = "this field is required"
		case "numeric":
			errFields[err.Field()] = "only numeric format"
		default:
			errFields[err.Field()] = err.Error()
		}
	}

	return errFields
}
