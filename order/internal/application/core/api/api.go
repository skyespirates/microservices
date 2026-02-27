package api

import (
	"context"
	"strings"
	"time"

	"github.com/skyespirates/microservices/order/internal/application/core/domain"
	"github.com/skyespirates/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {

	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	paymentErr := a.payment.Charge(ctx, &order)
	if paymentErr != nil {

		st := status.Convert(paymentErr)

		if st.Code() == codes.DeadlineExceeded {
			// Timeout case
			return domain.Order{}, status.Error(codes.DeadlineExceeded, "payment service did not respond within timeout")
		}

		var allErrors []string
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					allErrors = append(allErrors, violation.Description)
				}
			}
		}
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: strings.Join(allErrors, "\n"),
		}

		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)

		return domain.Order{}, statusWithDetails.Err()
	}

	return order, nil

}

func (a Application) GetOrder(orderID uint) (domain.Order, error) {
	order, err := a.db.Get(orderID)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
