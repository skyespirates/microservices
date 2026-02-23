package payment

import (
	"context"

	"github.com/skyespirates/microservices-proto/golang/payment"
	"github.com/skyespirates/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
	conn    *grpc.ClientConn
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := payment.NewPaymentClient(conn)

	return &Adapter{payment: client, conn: conn}, nil

}

func (a *Adapter) Charge(order *domain.Order) error {

	request := payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	}

	_, err := a.payment.Create(context.Background(), &request)
	return err
}

func (a *Adapter) Close() {
	a.conn.Close()
}
