package controller

import (
	"net/http"
)

func RegisterControllers() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	RegisterChatController()
}
