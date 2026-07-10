package bazars

import (
	"mms-dbsd/internal/domain/bazars/dto"
	"mms-dbsd/internal/domain/tenant"
	"mms-dbsd/internal/domain/users"

	"gorm.io/gorm"
)

type Bazars struct {
	gorm.Model
	BazarName        string `json:"bazar_name" gorm:"type:varchar(255);not null"`
	BazarDescription string `json:"bazar_description" gorm:"type:varchar(255);not null"`
	TenantCode       string `json:"tenant_code" gorm:"type:varchar(255);not null"`

	UserID uint `json:"user_id" gorm:"not null"`
	User   users.User `gorm:"foreignKey:UserID"`

	TenantID uint `json:"tenant_id"`
	Tenant   tenant.Tenant `gorm:"foreignKey:TenantID"`
}


func (b *Bazars) ToBazarResponse() *dto.BazarResponse {
	return &dto.BazarResponse{
		ID:               b.ID,
		BazarName:        b.BazarName,
		BazarDescription: b.BazarDescription,
		TenantCode:       b.TenantCode,
		CreatedAt:        b.CreatedAt,
	}
}