package main

import (
	"fmt"
	"nilchan/FinalProject/internal/models"
	datamainer "nilchan/FinalProject/internal/service"
	"nilchan/FinalProject/internal/transport/http"
)

func main() {

	fmt.Println("Start program")

	user := models.NewUser("Vanya")

	manager := datamainer.NewManagerMainer(user)
	handlers := http.NewHttpHandlers(manager, user)

	server := http.NewHttpServer(handlers)
	server.StartServer()

}
