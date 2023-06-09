package ports

import "github.com/bishtpramod19/microservices/order/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
