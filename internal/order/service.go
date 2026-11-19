package order

import "fmt"

type Order struct {
	Customer string
	Products []string
	Total    float64
	Status   string
}

type RepositoryWriter interface {
	Save(order Order) error
}

type Notifier interface {
	Send(customer string)
}

type OrderService struct {
	repository RepositoryWriter
	notifier   Notifier
}

func NewOrderService(repository RepositoryWriter, notifier Notifier) *OrderService {
	return &OrderService{repository: repository, notifier: notifier}
}

func (service *OrderService) CreateOrder(customer string, products []string, total float64) error {
	newOrder := Order{
		Customer: customer,
		Products: products,
		Total:    total,
		Status:   "pending",
	}
	if err := service.repository.Save(newOrder); err != nil {
		return fmt.Errorf("сохранить заказ: %w", err)
	}

	service.notifier.Send(customer)
	return nil
}
