package database

import (
	"database/sql"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) GetAll() ([]entity.Order, error) {
	stmt, err := r.Db.Prepare("SELECT * FROM orders")
	if err != nil {
		return []entity.Order{}, err
	}

	result, err := stmt.Query()
	if err != nil {
		return []entity.Order{}, err
	}

	var orders []entity.Order
	for result.Next() {
		var order entity.Order
		if err := result.Scan(
			&order.ID,
			&order.Price,
			&order.Tax,
			&order.FinalPrice); err != nil {
			return orders, err
		}
	}

	return orders, nil
}
