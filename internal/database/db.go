package database

import (
	"fmt"
	"mms-dbsd/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabse(cfg *config.Config) *gorm.DB {
	dsn := cfg.DBURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		fmt.Println("Server error:", err)
	}

	// auto generate table in database
	AutoMigrate(db)

	fmt.Println("Database connected successfully")

	return db
}