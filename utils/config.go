package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_SECRET string
	APP_ENV    string

	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_CALLBACK_URL  string
}

var config *Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config = &Config{
		APP_SECRET:           os.Getenv("APP_SECRET"),
		APP_ENV:              os.Getenv("APP_ENV"),
		GOOGLE_CLIENT_ID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GOOGLE_CLIENT_SECRET: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GOOGLE_CALLBACK_URL:  os.Getenv("GOOGLE_CALLBACK_URL"),
	}
}

func GetConfig() *Config {
	return config
}
