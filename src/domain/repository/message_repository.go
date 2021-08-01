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
	FindLast50Messages() ([]model.Message, error)
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
	data, jsonErr := json.Marshal(message)

	if jsonErr != nil {
		fmt.Println("Could not parse message to json")
	}

	repoErr := repo.client.RPush(key, data).Err()

	if repoErr != nil {
		fmt.Println("Could not save message", repoErr)
	}

	return message, repoErr
}

func (repo *repository) FindLast50Messages() ([]model.Message, error) {
	retrievedMessages, err := repo.client.LRange(key, -50, -1).Result()

	if err != nil {
		fmt.Println(err)
	}

	var messages []model.Message

	for _, message := range retrievedMessages {
		var msg model.Message

		if err := json.Unmarshal([]byte(message), &msg); err != nil {
			fmt.Println("Could not pare data to message")
		}

		messages = append(messages, msg)
	}

	return messages, err
}
