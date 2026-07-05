package tenant

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	TenantName string `json:"tenant_name" gorm:"type:varchar(150);not null"`
	Email      string `json:"email" gorm:"type:varchar(150);not null; unique"`
	SubDomain  string `json:"sub_domain" gorm:"type:varchar(50);not null;unique"`
}