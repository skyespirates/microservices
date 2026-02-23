package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/skyespirates/microservices-proto/golang/payment"
	"github.com/skyespirates/microservices/payment/config"
	"github.com/skyespirates/microservices/payment/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	payment.UnimplementedPaymentServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	log.Println("server running on port", a.port)

	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}
