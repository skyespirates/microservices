package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/skyespirates/microservices-proto/golang/order"
	"github.com/skyespirates/microservices/order/config"
	"github.com/skyespirates/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d, error: %w", a.port, err)
	}

	log.Println("gRPC server running on port", a.port)

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		return fmt.Errorf("failed to serve grpc on port %d, error: %w", a.port, err)
	}

	return nil

}
