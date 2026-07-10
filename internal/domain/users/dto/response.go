package dto

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"userName"`
	TenantID  uint    `json:"tenantId"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
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