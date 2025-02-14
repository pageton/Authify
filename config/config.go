package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabasePath string
	SecretKey    string
	Port         string
	LIMIT        int
	LoginPage    bool
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	limitStr := os.Getenv("LIMIT")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Fatalf("Invalid LIMIT value: %v", err)
	}

	loginPageStr := os.Getenv("LOGIN_PAGE")
	loginPage, err := strconv.ParseBool(loginPageStr)
	if err != nil {
		log.Printf("Invalid LOGIN_PAGE value, defaulting to false: %v", err)
		loginPage = false
	}

	return &Config{
		DatabasePath: os.Getenv("DATABASE_PATH"),
		SecretKey:    os.Getenv("SECRET_KEY"),
		Port:         os.Getenv("PORT"),
		LIMIT:        limit,
		LoginPage:    loginPage,
	}, nil
}
