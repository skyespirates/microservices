package ports

import "github.com/skyespirates/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id uint) (domain.Order, error)
	Save(*domain.Order) error
}
