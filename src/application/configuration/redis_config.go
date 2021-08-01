package configuration

import (
	"github.com/go-redis/redis"
	"log"
	"os"
)

func init() {
	LoadEnvironmentConfiguration()
}

func NewRedisClient() *redis.Client {
	url := os.Getenv("REDIS")

	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(opt)

	return client
}
