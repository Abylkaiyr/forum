package main

import (
	"log"

	"github.com/Abylkaiyr/forum/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	logger := apiserver.NewLogger()
	var api apiserver.APIServer
	s := api.NewServer(config, logger)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	s.Server()
}
