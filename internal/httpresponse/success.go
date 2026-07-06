package httpresponse

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(message string, data any) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}