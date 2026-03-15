package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetAllProduct() ([]models.Product, error) {

	rows, err := p.db.Query(context.Background(), `
		select id, product_name, product_desc, price, quantity, discount from products;
	`)

	products, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Product])

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) GetProductById(id int) (models.Product, error) {
	rows, err := p.db.Query(context.Background(), `
		select id, product_name, product_desc, price, quantity, discount from products where id = $1
	`, id)

	if err != nil {
		return models.Product{}, err
	}
	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])

	return result, err
}

func (p *ProductRepository) AddProduct(product models.Product) (models.Product, error) {
	query := `
		INSERT INTO products (product_name, product_desc, price, quantity, discount)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, product_name, product_desc, price, quantity, discount
	`
	rows, err := p.db.Query(context.Background(), query, product.Name, product.Description, product.Price, product.Quantity, product.Discount)
	if err != nil {
		return models.Product{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
}

func (p *ProductRepository) UpdateProductById(id int, product models.Product) (models.Product, error) {
	query := `
		UPDATE products 
		SET product_name = $1, product_desc = $2, price = $3, quantity = $4, discount = $5 
		WHERE id = $6 RETURNING id, product_name, product_desc, price, quantity, discount
	`
	rows, err := p.db.Query(context.Background(), query, product.Name, product.Description, product.Price, product.Quantity, product.Discount, id)
	if err != nil {
		return models.Product{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
}
