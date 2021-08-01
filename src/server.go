package main

import (
	"go-chat/src/application/configuration"
	"go-chat/src/application/controller"
	"go-chat/src/domain/service"
	"log"
	"net/http"
	"os"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	configuration.LoadEnvironmentConfiguration()
	controller.RegisterControllers()

	go service.MessageReceiver()

	log.Print("Server starting @ localhost:" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Couldn't start server")
	}
}
