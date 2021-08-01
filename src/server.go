package main

import (
	"go-chat/src/application/configuration"
	"go-chat/src/application/controller"
	"log"
	"net/http"
	"os"
)

func init() {
	configuration.LoadEnvironmentConfiguration()
}

var (
	port = os.Getenv("PORT")
)

func main() {
	controller.RegisterControllers()

	log.Print("Server starting @ localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Couldn't start server")
	}
}
