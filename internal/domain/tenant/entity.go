package tenant

import (
	"mms-dbsd/internal/domain/tenant/dto"

	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	TenantName string `json:"tenant_name" gorm:"type:varchar(150);not null"`
	Email      string `json:"email" gorm:"type:varchar(150);not null; unique"`
	SubDomain  string `json:"sub_domain" gorm:"type:varchar(50);not null;unique"`
}

func (t *Tenant) ToTenantResponse() *dto.TenantResponse {
	return &dto.TenantResponse{
		ID:        uint64(t.ID),
		TenantName: t.TenantName,
		Email:     t.Email,
		SubDomain: t.SubDomain,
		CreatedAt: t.CreatedAt,
	}
}