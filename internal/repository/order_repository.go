package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAllOrder() ([]models.Order, error) {
	query := `SELECT orders.id, orders.user_id, orders.status, orders.total, orders.image_path, orders.updated_at FROM orders;`

	rows, err := r.db.Query(context.Background(), query)

	orders, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Order])

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return orders, nil
}
