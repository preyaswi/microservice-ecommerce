package main

import (
	"log"

	"github.com/preyaswi/go-grpc-api-gateway/pkg/config"
	"github.com/preyaswi/go-grpc-api-gateway/pkg/di"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}
