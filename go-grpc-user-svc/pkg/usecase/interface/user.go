package interfaces

import (
	"github.com/preyaswi/go-grpc-user-svc/pkg/domain"
	"github.com/preyaswi/go-grpc-user-svc/pkg/utils/models"
)

type UserUseCase interface {
	UsersSignUp(user models.UserSignUp) (domain.TokenUser, error)
	UsersLogin(user models.UserLogin) (domain.TokenUser, error)
}
