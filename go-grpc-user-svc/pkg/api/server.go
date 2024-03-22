package server

import (
	"fmt"
	"net"

	"github.com/preyaswi/go-grpc-user-svc/pkg/config"
	"github.com/preyaswi/go-grpc-user-svc/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.UserServer) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	pb.RegisterUserServer(newServer,server)
	return &Server{
		server: newServer,
		listner: lis,
	},nil
}
func (c *Server)Start() error {
	fmt.Println("grpc server listening on port :50051")
    return c.server.Serve(c.listner)
}
