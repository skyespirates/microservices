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

func (a Adapter) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {

	o, err := a.api.GetOrder(uint(request.OrderId))
	if err != nil {
		return &order.GetOrderResponse{}, err
	}

	var oi []*order.OrderItem

	for _, v := range o.OrderItems {
		var o order.OrderItem
		o.ProductCode = v.ProductCode
		o.Quantity = v.Quantity
		o.UnitPrice = v.UnitPrice
		oi = append(oi, &o)
	}

	result := &order.GetOrderResponse{
		OrderId:    o.ID,
		OrderItems: oi,
	}

	return result, nil
}
