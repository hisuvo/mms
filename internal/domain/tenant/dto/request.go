package dto

type CreateTenantRequest struct {
	TenantName string `json:"tenant_name" validate:"required"`
	TenantCode string `json:"tenant_code" validate:"omitempty"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	SubDomain  string `json:"sub_domain" validate:"required"`
	IsActive   *bool  `json:"is_active" validate:"omitempty"`
}

type UpdateTenantRequest struct {
	TenantName string `json:"tenant_name" validate:"omitempty"`
	TenantCode string `json:"tenant_code" validate:"omitempty"`
	Email      string `json:"email" validate:"omitempty,email"`
	SubDomain  string `json:"sub_domain" validate:"omitempty"`
	IsActive   *bool  `json:"is_active" validate:"omitempty"`
}
