package main

import (
	"mms-dbsd/internal/config"
	"mms-dbsd/internal/database"
	"mms-dbsd/internal/server"
)

func main() {
	cfg := config.LoadEnv()

	db := database.ConnectDatabse(cfg)

	server.Start(db,cfg)
}