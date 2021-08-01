package configuration

import (
	"github.com/go-redis/redis"
	"log"
	"os"
)

func NewRedisClient() *redis.Client {
	url := os.Getenv("REDIS")

	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(opt)

	_, redisError := client.Ping().Result()
	if redisError != nil {
		log.Fatal("Could not connect to redis")
	}

	return client
}
