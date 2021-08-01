package reposiroty

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"go-chat/src/domain/model"
	"go-chat/src/domain/repository"
	"log"
	"os"
	"testing"
)

var (
	client *redis.Client
)

const key = "chat"

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()

	if err != nil {
		log.Fatalf("Error while creating miniredis instance: %s", err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	code := m.Run()
	os.Exit(code)
}

func TestWhenSavingMessageIfRepositoryDoesNotReturnErrorAndSavesMessage(t *testing.T) {
	client.Del(key)

	repo := repository.NewMessageRepository(client)

	message, err := repo.Save(mockMessage())

	assert.Equal(t, nil, err)

	assert.ObjectsAreEqual(mockMessage(), message)
	assert.Equal(t, mockMessage(), message)
}

func TestWhenFindingLast50MessagesRepositoryReturnsUpTo50Messages(t *testing.T) {
	client.Del(key)

	firstRun := 25
	secondRun := 80

	expected := 50

	repo := repository.NewMessageRepository(client)

	for i := 0; i < firstRun; i++ {
		repo.Save(mockMessage())
	}

	messages := repo.FindLast50Messages()
	assert.Equal(t, firstRun, len(messages))

	for i := 0; i < secondRun; i++ {
		repo.Save(mockMessage())
	}

	messages = repo.FindLast50Messages()
	assert.Equal(t, expected, len(messages))
}

func mockMessage() model.Message {
	var message = model.Message{
		Username: "Worker",
		Content:  "work work!",
	}

	return message
}
