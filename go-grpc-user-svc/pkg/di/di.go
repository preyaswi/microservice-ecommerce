package di

import (
	server "github.com/preyaswi/go-grpc-user-svc/pkg/api"
	"github.com/preyaswi/go-grpc-user-svc/pkg/api/service"
	"github.com/preyaswi/go-grpc-user-svc/pkg/config"
	"github.com/preyaswi/go-grpc-user-svc/pkg/db"
	"github.com/preyaswi/go-grpc-user-svc/pkg/repository"
	"github.com/preyaswi/go-grpc-user-svc/pkg/usecase"
)
func InitializeApi(cfg config.Config) (*server.Server,error) {
	gormDB, err := db.ConnectToDb(cfg)
	if err != nil {
		return nil, err
	}
	adminRepository:=repository.NewUserRepository(gormDB)
	adminUseCase:=usecase.NewUserUseCase(adminRepository)

	userServiceServer:=service.NewUserServer(adminUseCase)
	grpcServer,err:=server.NewGRPCServer(cfg,userServiceServer)
	if err!=nil{
		return &server.Server{}, err
	}
	return grpcServer,nil
}