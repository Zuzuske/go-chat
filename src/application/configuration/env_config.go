package configuration

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironmentConfiguration() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment file")
	}
}
