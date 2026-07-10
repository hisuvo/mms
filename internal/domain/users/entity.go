package users

import (
	"mms-dbsd/internal/domain/tenant"
	"mms-dbsd/internal/domain/users/dto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"size:100;not null"`
	TenantID uint `json:"tenant_id" gorm:"size:100;not null;index"`
	Tenant   tenant.Tenant `gorm:"foreignKey:TenantID"`
	Phone    string `json:"phone" gorm:"size:20;uniqueIndex"`
	Email    string `json:"email" gorm:"size:255;not null;uniqueIndex"`
	Password string `json:"password" gorm:"size:255;not null"`
	Role     string `json:"role" gorm:"size:50;not null;default:user"`
}

func (u *User) ToUserResponse() *dto.UserResponse {
	return &dto.UserResponse{
		ID:        u.ID,
		UserName:  u.UserName,
		TenantID:  u.TenantID,
		Phone:     u.Phone,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}