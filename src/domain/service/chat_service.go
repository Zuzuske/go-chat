package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go-chat/src/application/configuration"
	"go-chat/src/domain/model"
	"go-chat/src/domain/repository"
	"net/http"
)

var (
	messageRepository = repository.NewMessageRepository(configuration.NewRedisClient())
	clients           = make(map[*websocket.Conn]bool)
	broadcaster       = make(chan model.Message)
	upgrader          = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

func UpgradeToWSConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Could not upgrade to websocket")
		return
	}

	defer func(ws *websocket.Conn) {
		if err := ws.Close(); err != nil {
			fmt.Println(err)
		}
	}(conn)

	clients[conn] = true

	loadChatHistory(conn)

	messageSender(conn)
}

func messageSender(conn *websocket.Conn) {
	for {
		var message model.Message

		if err := conn.ReadJSON(&message); err != nil {
			delete(clients, conn)
			break
		}

		broadcaster <- message
	}
}

func loadChatHistory(conn *websocket.Conn) {
	messages, err := messageRepository.FindLast50Messages()

	if err != nil {
		return
	}

	for _, message := range messages {
		sendMessage(conn, message)
	}
}

func MessageReceiver() {
	for {
		message := <-broadcaster

		saveMessage(message)
		sendMessageToClients(message)
	}
}

func saveMessage(message model.Message) {
	_, err := messageRepository.Save(message)
	if err != nil {
		return
	}
}

func sendMessageToClients(message model.Message) {
	for client := range clients {
		sendMessage(client, message)
	}
}

func sendMessage(client *websocket.Conn, message model.Message) {
	err := client.WriteJSON(message)

	if err != nil {
		fmt.Println(err)

		if err := client.Close(); err != nil {
			fmt.Println(err)
		}

		delete(clients, client)
	}
}
