package tenant

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	TenantName string `json:"tenant_name" gorm:"not null"`
	Email      string `json:"email" gorm:"not null"`
}