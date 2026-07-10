package tenant

import (
	"mms-dbsd/internal/domain/tenant/dto"

	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	TenantName string `json:"tenant_name" gorm:"type:varchar(150);not null;unique"`
	TenantCode string `json:"tenant_code" gorm:"type:varchar(20);not null;unique"`
	Email      string `json:"email" gorm:"type:varchar(150);not null; unique"`
	Password   string `json:"password" gorm:"type:varchar(255);not null"`
	SubDomain  string `json:"sub_domain" gorm:"type:varchar(50);not null;unique"`
	IsActive   bool   `json:"is_active" gorm:"type:boolean;not null;default:true"`
}

func (t *Tenant) ToTenantResponse() *dto.TenantResponse {
	return &dto.TenantResponse{
		ID:         uint64(t.ID),
		TenantName: t.TenantName,
		TenantCode: t.TenantCode,
		Email:      t.Email,
		SubDomain:  t.SubDomain,
		IsActive:   t.IsActive,
		CreatedAt:  t.CreatedAt,
	}
}