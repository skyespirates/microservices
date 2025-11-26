package db

import (
	"fmt"

	"github.com/skyespirates/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	dst := []interface{}{&Order{}, &OrderItem{}}

	err := db.AutoMigrate(dst...)
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db}, nil
}

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)

	var orderItems []domain.OrderItem

	for _, oi := range orderEntity.OrderItems {
		var orderItem domain.OrderItem
		orderItem.ProductCode = oi.ProductCode
		orderItem.UnitPrice = oi.UnitPrice
		orderItem.Quantity = oi.Quantity

		orderItems = append(orderItems, orderItem)
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, oi := range order.OrderItems {
		var orderItem OrderItem
		orderItem.ProductCode = oi.ProductCode
		orderItem.UnitPrice = oi.UnitPrice
		orderItem.Quantity = oi.Quantity

		orderItems = append(orderItems, orderItem)
	}

	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	res := a.db.Create(&orderModel)
	if res == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}
