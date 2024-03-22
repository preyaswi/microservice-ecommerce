package interfaces

import (
	"github.com/preyaswi/go-grpc-admin-svc/pkg/domain"
	"github.com/preyaswi/go-grpc-admin-svc/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
}
