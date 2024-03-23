package main

import (
	"log"

	"github.com/preyaswi/grpc-cart-svc/pkg/config"
	"github.com/preyaswi/grpc-cart-svc/pkg/di"
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
