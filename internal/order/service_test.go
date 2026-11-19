package order_test

import (
	"errors"
	"testing"

	"example/solid/internal/order"
)

type failingRepository struct {
	err error
}

func (repo failingRepository) Save(order.Order) error {
	return repo.err
}

type countingNotifier struct {
	calls int
}

func (notifier *countingNotifier) Send(string) {
	notifier.calls++
}

func TestCreateOrderSkipsNotificationOnSaveError(t *testing.T) {
	saveError := errors.New("запись недоступна")
	notifier := &countingNotifier{}
	service := order.NewOrderService(failingRepository{err: saveError}, notifier)

	if err := service.CreateOrder("Иван", []string{"apple"}, 10.5); !errors.Is(err, saveError) {
		t.Fatalf("CreateOrder error = %v, want %v", err, saveError)
	}
	if notifier.calls != 0 {
		t.Errorf("Send called %d times, want 0", notifier.calls)
	}
}
