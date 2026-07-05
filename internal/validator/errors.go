package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationError(err error) []ErrorDetail {
	var details []ErrorDetail

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			details = append(details, ErrorDetail{
				Field:   e.Field(), // or use JSON tag
				Message: validationMessage(e),
			})
		}
	}

	return details
}

func validationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return "Invalid email address"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	default:
		return "Invalid value"
	}
}