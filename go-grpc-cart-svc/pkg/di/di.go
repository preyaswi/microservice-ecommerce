package di

import (
	server "github.com/preyaswi/grpc-cart-svc/pkg/api"
	"github.com/preyaswi/grpc-cart-svc/pkg/api/service"
	"github.com/preyaswi/grpc-cart-svc/pkg/client"
	"github.com/preyaswi/grpc-cart-svc/pkg/config"
	"github.com/preyaswi/grpc-cart-svc/pkg/db"
	"github.com/preyaswi/grpc-cart-svc/pkg/repository"
	"github.com/preyaswi/grpc-cart-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	cartRepository := repository.NewCartRepository(gormDB)
	productClient := client.NewProductClient(&cfg)
	adminUseCase := usecase.NewCartUseCase(cartRepository, productClient)

	adminServiceServer := service.NewCartServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
