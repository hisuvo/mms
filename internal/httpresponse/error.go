package httpresponse

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func NewErrorResponse(message string, errs any) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errs,
	}
}

func ValidationError(errs any) *ErrorResponse {
	return NewErrorResponse(
		"Validation failed",
		errs,
	)
}

func BadRequest(message string) *ErrorResponse {
	return NewErrorResponse(
		message,
		nil,
	)
}

func Unauthorized(message string) *ErrorResponse {
	return NewErrorResponse(
		message,
		nil,
	)
}

func Forbidden(message string) *ErrorResponse {
	return NewErrorResponse(
		message,
		nil,
	)
}

func NotFound(message string) *ErrorResponse {
	return NewErrorResponse(
		message,
		nil,
	)
}

func Conflict(message string) *ErrorResponse {
	return NewErrorResponse(
		message,
		nil,
	)
}

func TooManyRequests(message string) *ErrorResponse {
	return NewErrorResponse(
		message,
		nil,
	)
}

func InternalServerError() *ErrorResponse {
	return NewErrorResponse(
		"Internal server error",
		nil,
	)
}