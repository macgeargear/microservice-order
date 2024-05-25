package ports

import "github.com/macgeargear/microservices-order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
