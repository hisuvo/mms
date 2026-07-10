package database

import (
	"log"

	"mms-dbsd/internal/domain/tenant"
	"mms-dbsd/internal/domain/users"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&users.User{},
		&tenant.Tenant{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully")
}