package main

import (
	"database/sql"
	"fmt"
	"log"

	"example/solid/internal/notify"
	"example/solid/internal/order"
	"example/solid/internal/sqlite"

	_ "modernc.org/sqlite"
)

type repositoryInitializer interface {
	Init() error
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.Open("sqlite", "orders.db")
	if err != nil {
		return fmt.Errorf("открыть SQLite: %w", err)
	}
	defer db.Close()

	repository := sqlite.NewSQLiteRepo(db)
	var initializer repositoryInitializer = repository
	if err := initializer.Init(); err != nil {
		return fmt.Errorf("создать таблицу заказов: %w", err)
	}

	emailOrders := order.NewOrderService(repository, notify.EmailSender{})
	if err := emailOrders.CreateOrder("Иван", []string{"apple", "banana"}, 10.5); err != nil {
		return err
	}

	smsOrders := order.NewOrderService(repository, notify.SMSSender{})
	return smsOrders.CreateOrder("Мария", []string{"pear"}, 7.25)
}
