package sqlite

import (
	"database/sql"
	"fmt"

	"example/solid/internal/order"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

func (repo *SQLiteRepo) Init() error {
	_, err := repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customer TEXT NOT NULL,
			products TEXT NOT NULL,
			total REAL NOT NULL,
			status TEXT NOT NULL
		)
	`)
	return err
}

func (repo *SQLiteRepo) Save(newOrder order.Order) error {
	_, err := repo.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		newOrder.Customer, fmt.Sprint(newOrder.Products), newOrder.Total, newOrder.Status,
	)
	return err
}
