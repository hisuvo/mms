package dto

type BazarCreateRequest struct {
	BazarName        string `json:"bazar_name" validate:"required"`
	BazarDescription string `json:"bazar_description" validate:"required"`
	TenantCode       string `json:"tenant_code" validate:"required"`
}

type BazarUpdateRequest struct {
	BazarName        string `json:"bazar_name" validate:"omitempty"`
	BazarDescription string `json:"bazar_description" validate:"omitempty"`
	TenantCode       string `json:"tenant_code" validate:"omitempty"`
}