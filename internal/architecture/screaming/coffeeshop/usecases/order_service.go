package usecases

import (
	"log"

	"github.com/qjcg/examples-go/internal/architecture/screaming/coffeeshop/entities"
	"github.com/qjcg/examples-go/internal/architecture/screaming/coffeeshop/infra"
)

type OrderService struct {
	Store infra.OrderStore
}

func NewOrderService(store infra.OrderStore) *OrderService {
	return &OrderService{Store: store}
}

func (s *OrderService) AddItem(items ...entities.Item) {
	order := entities.Order{
		Items: items,
	}

	s.Store.Save(order)
}

func (s *OrderService) Total() float64 {
	var total float64

	orders, err := s.Store.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(orders); i++ {
		total += orders[i].Price
	}

	return total
}
