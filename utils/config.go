package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Foo string
}

var config *Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config = &Config{
		Foo: os.Getenv("FOO"),
	}
}

func GetConfig() *Config {
	return config
}
