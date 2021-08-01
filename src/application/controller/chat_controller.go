package controller

import (
	"go-chat/src/application/facade"
	"net/http"
)

func RegisterChatController()  {
	http.HandleFunc("/ws", facade.UpgradeToWSConnection)
}