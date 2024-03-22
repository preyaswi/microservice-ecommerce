package interfaces

import (
	"github.com/preyaswi/go-grpc-user-svc/pkg/domain"
	"github.com/preyaswi/go-grpc-user-svc/pkg/utils/models"
)

type UserRepository interface {
	CheckUserExistsByEmail(email string) (*domain.User, error)
	CheckUserExistsByPhone(phone string) (*domain.User, error)
	UserSignUp(user models.UserSignUp) (models.UserDetails, error)
	FindUserByEmail(user models.UserLogin) (models.UserDetail, error)
}
