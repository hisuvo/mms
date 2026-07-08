package database

import (
	"fmt"
	"log"
	"mms-dbsd/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabse(cfg *config.Config) *gorm.DB {
	dsn := cfg.DBURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Get underlying *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(100)               // Maximum open connections
	sqlDB.SetMaxIdleConns(20)                // Maximum idle connections
	sqlDB.SetConnMaxLifetime(time.Hour)      // Recreate connection after 1 hour
	sqlDB.SetConnMaxIdleTime(15 * time.Minute) // Close idle connections after 15 minutes

	// auto generate table in database
	AutoMigrate(db)

	// Super Admin seed function call here
	if err := Seed(db, cfg); err == nil {
		fmt.Println("Super Admin seed completed successfully")
	}else{
		fmt.Println("Super Admin seed failed")
	}

	fmt.Println("Database connected successfully")

	return db
}