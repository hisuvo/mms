package dto

import "time"

type TenantResponse struct {
	ID         uint64    `json:"id"`
	TenantName string    `json:"tenant_name"`
	TenantCode string    `json:"tenant_code"`
	Email      string    `json:"email"`
	SubDomain  string    `json:"sub_domain"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}

type TenantListResponse struct {
	Tenants []TenantResponse `json:"tenants"`
	Total   int64            `json:"total"`
}
