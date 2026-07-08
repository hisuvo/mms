package database

import (
	"mms-dbsd/internal/config"
	"mms-dbsd/internal/domain/users"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB, cfg *config.Config) error {
	var count int64
	db.Model(&users.User{}).Count(&count)

	if count > 0 {
		return nil
	}

	user := users.User{
		UserName: cfg.SuperAdminName,
		Email:    cfg.SuperAdminEmail,
		Password: cfg.SuperAdminPass,
		Phone:    cfg.SuperAdminPhone,
		Role:     cfg.SuperAdminRole,
	}

	return db.Create(&user).Error
}

// Note: lange project seed file structer
/*

	// 1) seed file structer
	internal/
	└── database/
		├── seed.go
		└── seeds/
			├── tenant_seed.go
			├── user_seed.go
			├── meal_seed.go
			└── role_seed.go

	// 2) run seed file
		func Seed(db *gorm.DB) error {
		if err := SeedRoles(db); err != nil {
			return err
		}
		if err := SeedTenants(db); err != nil {
			return err
		}
		if err := SeedUsers(db); err != nil {
			return err
		}
		return nil
	}
*/

