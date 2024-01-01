package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Foo string

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
		Foo:                  os.Getenv("FOO"),
		GOOGLE_CLIENT_ID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GOOGLE_CLIENT_SECRET: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GOOGLE_CALLBACK_URL:  os.Getenv("GOOGLE_CALLBACK_URL"),
	}
}

func GetConfig() *Config {
	return config
}
