package main

import (
	"log"

	"github.com/madxiii/axxonsofttask/http"
	"github.com/madxiii/axxonsofttask/repository"
	"github.com/madxiii/axxonsofttask/server"
	"github.com/madxiii/axxonsofttask/service"
)

const address string = ":8888"

func main() {
	localStore := repository.NewHashMap()

	repos := repository.New(localStore)
	services := service.New(*repos)
	handler := http.NewHandler(*services).Routes()

	serv := new(server.Server)
	if err := serv.Run(address, handler); err != nil {
		log.Fatalf("error occured while running server: %s\n", err.Error())
	}
}
