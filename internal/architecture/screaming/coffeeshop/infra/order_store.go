package infra

import (
	"github.com/qjcg/examples-go/internal/architecture/screaming/coffeeshop/entities"
)

type OrderStore interface {
	Save(order entities.Order) error
	FindAll() ([]entities.Order, error)
}
