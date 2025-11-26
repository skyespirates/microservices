package grpc

import "github.com/skyespirates/microservices/order/internal/ports"

type Adapter struct {
	api  ports.APIPort
	port int
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api, port}
}

func (a Adapter) run() {}
