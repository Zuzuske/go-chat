package service

import (
	"go-chat/src/domain/model"
)



func mockMessage() model.Message {
	var message = model.Message{
		Username: "Worker",
		Content:  "work work!",
	}

	return message
}
