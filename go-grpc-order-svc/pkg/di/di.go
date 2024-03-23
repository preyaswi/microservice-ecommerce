package di

import (
	server "github.com/preyaswi/go-grpc-order-svc/pkg/api"
	"github.com/preyaswi/go-grpc-order-svc/pkg/api/client"
	"github.com/preyaswi/go-grpc-order-svc/pkg/api/service"
	"github.com/preyaswi/go-grpc-order-svc/pkg/config"
	"github.com/preyaswi/go-grpc-order-svc/pkg/db"
	"github.com/preyaswi/go-grpc-order-svc/pkg/repository"
	"github.com/preyaswi/go-grpc-order-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	orderRepository := repository.NewOrderRepository(gormDB)
	cartClient := client.NewCartClient(&cfg)
	productClient := client.NewProductClient(&cfg)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartClient, productClient)

	orderServiceServer := service.NewOrderServer(orderUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, orderServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
