run:
		go run cmd/main.go

proto:
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/cart/cart.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/order/order.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/product/product.proto