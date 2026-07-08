package dto

type RegisterRequest struct {
	UserName string `json:"userName" validate:"required,min=3,max=50"`
	TenantID string `json:"tenantID" validate:"required"`
	Phone    string `json:"phone" validate:"required,numeric,min=10,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type UpdateRequest struct {
	UserName string `json:"userName" validate:"omitempty,min=3,max=50"`
	Phone    string `json:"phone" validate:"omitempty,numeric,min=10,max=15"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Role     string `json:"role" validate:"omitempty,oneof=admin user"`
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}