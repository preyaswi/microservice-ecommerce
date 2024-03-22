package di

import (
	server "github.com/preyaswi/go-grpc-product-svc/pkg/api"
	"github.com/preyaswi/go-grpc-product-svc/pkg/api/service"
	"github.com/preyaswi/go-grpc-product-svc/pkg/config"
	"github.com/preyaswi/go-grpc-product-svc/pkg/db"
	"github.com/preyaswi/go-grpc-product-svc/pkg/repository"
	"github.com/preyaswi/go-grpc-product-svc/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewProductRepository(gormDB)
	adminUseCase := usecase.NewProductUseCase(adminRepository)

	adminServiceServer := service.NewProductServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
