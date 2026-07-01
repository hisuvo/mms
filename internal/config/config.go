package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT           string
	DBURL          string
	JWTSecret      string
	SecretDuration string
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		fmt.Printf("%s is required", key)
	}
	return value
}

func LoadEnv() *Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not loaded")
	}

	return &Config{
		PORT:           mustGetEnv("PORT"),
		DBURL:          mustGetEnv("DB_URL"),
		JWTSecret:      mustGetEnv("JWT_SECRET"),
		SecretDuration: mustGetEnv("SECRET_DURATION"),
	}
}