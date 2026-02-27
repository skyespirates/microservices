package ports

import (
	"context"

	"github.com/skyespirates/microservices/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(orderID uint) (domain.Order, error)
}
