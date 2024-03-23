package interfaces

import (
	"github.com/preyaswi/go-grpc-order-svc/pkg/domain"
	"github.com/preyaswi/go-grpc-order-svc/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error)
}
