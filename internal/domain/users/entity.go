package users

import (
	"mms-dbsd/internal/domain/users/dto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"size:100;not null"`
	TenantID string `gorm:"size:100;not null;index"`
	Phone    string `gorm:"size:20;uniqueIndex"`
	Email    string `gorm:"size:255;not null;uniqueIndex"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"size:50;not null;default:user"`
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