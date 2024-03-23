package client

import (
	"context"
	"fmt"

	pb "github.com/preyaswi/go-grpc-order-svc/pkg/pb/product"
	"github.com/preyaswi/go-grpc-order-svc/pkg/config"
	"google.golang.org/grpc"
)

type clientProduct struct {
	Client pb.ProductClient
}

func NewProductClient(cfg *config.Config) *clientProduct {

	grpcConnection, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductClient(grpcConnection)

	return &clientProduct{
		Client: grpcClient,
	}

}
func (c *clientProduct) ProductStockMinus(productID, stock int) error {
	_, err := c.Client.ProductStockMinus(context.Background(), &pb.ProductStockMinusRequest{
		ID:    int64(productID),
		Stock: int64(stock),
	})
	if err != nil {
		return err
	}
	return nil
}
