package repository

import (
	"backend/internal/models"
	"context"
	"errors"

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
		VALUES ($1, $2, $3, $4, $5)
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

func (p *ProductRepository) DeleteProductById(id int) error {
	query := `DELETE FROM products WHERE id = $1`

	result, err := p.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("no product found with this id")
	}

	return nil
}

// repository card landingpage

func (p *ProductRepository) GetAllProductHome() ([]models.ProductHome, error) {

	rows, err := p.db.Query(context.Background(), `
		SELECT
			p.id,
			p.product_name,
			p.product_desc,
			p.price,
			product_images.path,
			reviews.ratings
		FROM products p
		LEFT JOIN product_images ON p.id = product_images.product_images_id
		LEFT JOIN reviews ON p.id = reviews.review_id
		WHERE p.id > 5
		LIMIT 4;
	`)

	products, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.ProductHome])

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) ProductReview() ([]models.ReviewProduct, error) {

	rows, err := p.db.Query(context.Background(), `
		SELECT 
			users.full_name,
			users.pictures,
			reviews.messages,
			reviews.ratings
		FROM users
		LEFT JOIN reviews ON users.id = reviews.review_id
		WHERE reviews.ratings = (SELECT MAX(ratings) FROM reviews);
	`)

	review, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.ReviewProduct])

	if err != nil {
		return nil, err
	}
	return review, nil
}

// repository browse product
