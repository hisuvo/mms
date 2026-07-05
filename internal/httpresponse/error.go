package httpresponse

import (
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func NewErrorResponse(code int, message string, errors any) *Error {
	return &Error{
		Code:    code,
		Success: false,
		Message: message,
		Errors:  errors,
	}
}

func ValidationError(errors any) *Error {
	return NewErrorResponse(http.StatusUnprocessableEntity, "Validation failed", errors)
}

func BadRequest(message string) *Error {
	return NewErrorResponse(http.StatusBadRequest, message, nil)
}

func Unauthorized(message string) *Error {
	return NewErrorResponse(http.StatusUnauthorized, message, nil)
}

func Conflict(message string) *Error {
	return NewErrorResponse(http.StatusConflict, message, nil)
}

func InternalServerError() *Error {
	return NewErrorResponse(http.StatusInternalServerError, "Internal server error", nil)
}