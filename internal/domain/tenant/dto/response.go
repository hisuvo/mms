package dto

type TenantResponse struct {
	ID        string `json:"id"`
	TenatName string `json:"tenat_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type TenantListResponse struct {
	Tenants []TenantResponse `json:"tenants"`
	Total   int64            `json:"total"`
}
