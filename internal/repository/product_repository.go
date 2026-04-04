package repository

import (
	"backend/internal/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
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
	defer rows.Close()

	return products, nil
}

func (p *ProductRepository) GetProductById(id int) (models.Product, error) {
	rows, err := p.db.Query(context.Background(), `
		select id, product_name, product_desc, price, quantity, discount from products where id = $1
	`, id)

	if err != nil {
		return models.Product{}, err
	}
	defer rows.Close()

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

	defer rows.Close()

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

	defer rows.Close()

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

func (p *ProductRepository) GetAllProductHome(ctx context.Context) ([]models.ProductHome, error) {

	rows, err := p.db.Query(ctx, `
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
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.ProductHome])

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) ProductReview(ctx context.Context) ([]models.ReviewProduct, error) {

	query := `
	SELECT
		users.full_name,
		users.pictures,
		reviews.messages,
		reviews.ratings
	FROM reviews
	INNER JOIN users ON users.id = reviews.user_id
	WHERE reviews.ratings >= 4
	LIMIT 3;
	`

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	review, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.ReviewProduct])

	if err != nil {
		return nil, err
	}
	return review, nil
}

// repository browse product
func (p *ProductRepository) BrowseProducts() ([]models.BrowseProduct, error) {

	rows, err := p.db.Query(context.Background(), `
		SELECT p.id, p.product_name, p.product_desc, p.price, p.quantity, p.discount, d.is_flash_sale, COALESCE(ARRAY_AGG(pi.path) FILTER (WHERE pi.path IS NOT NULL), '{}') as images
		FROM products p
		LEFT JOIN discount d ON p.id = d.discount_id 
		LEFT JOIN product_images pi ON p.id = pi.product_id        
		GROUP BY p.id, d.discount_id, d.is_flash_sale
		ORDER BY id asc
	`)

	products, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.BrowseProduct])

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return products, nil
}

func (p *ProductRepository) DetailProduct(id int) ([]models.DetailProduct, error) {

	rows, err := p.db.Query(context.Background(), `
	SELECT
		p.*,

		d.discount_rate,
		d.is_flash_sale,

		(SELECT JSON_AGG(path)
		FROM product_images
		WHERE product_id = p.id) AS images,

		(SELECT JSON_AGG(
			JSON_BUILD_OBJECT(
				'size_name', size_name,
				'size_price', size_price
			)
		)
		FROM product_size
		WHERE product_id = p.id) AS sizes,

		(SELECT JSON_AGG(
			JSON_BUILD_OBJECT(
				'variant_name', variant_name,
				'add_price', add_price
			)
		)
		FROM product_variant
		WHERE product_id = p.id) AS variants,

		(SELECT AVG(ratings)
		FROM reviews
		WHERE product_id = p.id) AS rating

	FROM products p
	LEFT JOIN discount d ON p.id = d.discount_id
	WHERE p.id = $1;
	`, id)

	products, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.DetailProduct])

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return products, nil
}
