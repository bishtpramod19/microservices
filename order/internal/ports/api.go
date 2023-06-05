package ports

import "github.com/bishtpramod19/microservices/order/order/internal/application/core/domain"

type APIPort interface {
	placeOrder(order domain.Order) (domain.Order, error)
}
