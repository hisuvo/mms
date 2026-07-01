package users

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserName string `gorm:"size:100;not null"`
	TenantID string `gorm:"size:100;not null;index"`
	Phone    string `gorm:"size:20;uniqueIndex"`
	Email    string `gorm:"size:255;not null;uniqueIndex"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"size:50;not null;default:user"`

	CreatedAt time.Time
	UpdatedAt time.Time
}