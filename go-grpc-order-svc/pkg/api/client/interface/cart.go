package interfaces

import "github.com/preyaswi/go-grpc-order-svc/pkg/utils/models"

type CartClient interface {
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
	DoesCartExist(userID int) (bool, error)
	TotalAmountInCart(userID int) (float64, error)
	UpdateCartAfterOrder(userID, productID int, quantity float64) error
}
