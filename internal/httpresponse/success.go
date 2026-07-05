package httpresponse

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(response *SuccessResponse) *SuccessResponse {
	response.Success = true
	return response
}