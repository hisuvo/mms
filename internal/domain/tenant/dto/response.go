package dto

import "time"

type TenantResponse struct {
	ID         uint64    `json:"id"`
	TenantName string    `json:"tenant_name"`
	Email      string    `json:"email"`
	SubDomain  string    `json:"sub_domain"`
	CreatedAt  time.Time `json:"created_at"`
}

type TenantListResponse struct {
	Tenants []TenantResponse `json:"tenants"`
	Total   int64            `json:"total"`
}
