package dto

type RegisterRequest struct {
	UserName string `json:"user_name" validate:"required,min=3,max=50"`
	TenantID string `json:"tenant_id" validate:"required"`
	Phone    string `json:"phone" validate:"required,numeric,min=10,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type UpdateRequest struct {
}