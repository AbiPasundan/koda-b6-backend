package repository

import (
	"backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) GetAllOrder(conn *pgx.Conn) []models.Order {
	test := []models.Order{}
	return test
}
