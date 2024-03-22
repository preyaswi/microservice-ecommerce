package interfaces

import (
	"github.com/preyaswi/go-grpc-admin-svc/pkg/domain"
	"github.com/preyaswi/go-grpc-admin-svc/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
}
