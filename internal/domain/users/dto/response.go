package dto

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name"`
	TenantID  uint      `json:"tenant_id"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	User         UserResponse `json:"user"`
	Token        TokenResponse `json:"token"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}