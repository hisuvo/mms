package main

import (
	"fmt"
	"mms-dbsd/internal/auth"
	"mms-dbsd/internal/config"
	"mms-dbsd/internal/database"
	"mms-dbsd/internal/server"
	"time"
)

func main() {
	cfg := config.LoadEnv()

	db := database.ConnectDatabse(cfg)

	jwtService := auth.NewJWT([]byte(cfg.JWTSecret), 24*time.Hour)

	fmt.Println("jwtToken",jwtService)
	
	server.Start(db,cfg)
}