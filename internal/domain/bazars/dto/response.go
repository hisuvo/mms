package dto

import "time"

type BazarResponse struct {
	ID               uint `json:"id"`
	BazarName        string `json:"bazar_name"`
	BazarDescription string `json:"bazar_description"`
	TenantCode       string `json:"tenant_code"`
	CreatedAt        time.Time `json:"created_at"`
}