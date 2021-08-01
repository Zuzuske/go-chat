package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"go-chat/src/domain/model"
)

const key = "chat"

type MessageRepository interface {
	Save(message model.Message) (model.Message, error)
	FindLast50Messages() []string
}

type repository struct {
	client *redis.Client
}

func NewMessageRepository(redisClient *redis.Client) MessageRepository {
	return &repository{
		client: redisClient,
	}
}

func (repo *repository) Save(message model.Message) (model.Message, error) {
	data, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Could not parse message to json")
	}

	repoErr := repo.client.RPush(key, data).Err();

	if err != nil {
		fmt.Println("Could not save message", repoErr)
	}

	return message, repoErr
}

func (repo *repository) FindLast50Messages() []string {
	messages, err := repo.client.LRange(key, -50, -1).Result()

	if err != nil {
		fmt.Println(err)
	}

	return messages
}
