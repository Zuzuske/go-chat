package configuration

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	LoadEnvironmentConfiguration()
}

func LoadEnvironmentConfiguration() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment file")
	}
}
