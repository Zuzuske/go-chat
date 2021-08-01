package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	LoadEnvironmentConfiguration()
}

func LoadEnvironmentConfiguration() {
	fmt.Println("config")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment file")
	}
}
