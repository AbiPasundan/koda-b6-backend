package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetAllProduct(conn *pgx.Conn) ([]models.Product, error) {
	rows, err := conn.Query(context.Background(), `
		select id, product_name, product_desc, price, quantity, product_category, discount from products;
	`)

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])

	if err != nil {
		return nil, err
	}
	return products, nil
}
