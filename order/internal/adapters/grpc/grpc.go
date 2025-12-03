package grpc

import (
	"context"

	"github.com/skyespirates/microservices-proto/golang/order"
	"github.com/skyespirates/microservices/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem

	for _, oi := range request.OrderItems {
		var orderItem domain.OrderItem
		orderItem.ProductCode = oi.ProductCode
		orderItem.UnitPrice = oi.UnitPrice
		orderItem.Quantity = oi.Quantity

		orderItems = append(orderItems, orderItem)
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
