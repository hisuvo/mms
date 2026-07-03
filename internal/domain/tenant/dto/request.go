package dto

type CreateTenantRequest struct {
	TenantName string `json:"tenant_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
}

type UpdateTenantRequest struct {
	TenantName *string `json:"tenant_name" validate:"omitempty"`
	Email      *string `json:"email" validate:"omitempty,email"`
}
