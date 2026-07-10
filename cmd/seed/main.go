package main

import (
	"fmt"
	"log"
	"mms-dbsd/internal/config"
	"mms-dbsd/internal/database"
)

func main() {

	cfg := config.LoadEnv()

	db := database.ConnectDatabse(cfg)

	err:= database.Seed(db, cfg)
	if err != nil {
		log.Fatal("Seed failed:", err)
	}

	fmt.Println("Seed completed successfully")

}