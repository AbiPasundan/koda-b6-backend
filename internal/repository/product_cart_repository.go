package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductCartRepository struct {
	db *pgxpool.Pool
}

func NewProductCartRepository(db *pgxpool.Pool) *ProductCartRepository {
	return &ProductCartRepository{db: db}
}

func (r *ProductCartRepository) GetCart(id int) ([]models.ProductCart, error) {
	query := `
		SELECT 
			ci.cart_item_id,
			ci.product_id,
			ci.product_name,
			ci.variant_name,
			ci.size_name,
			ci.base_price,
			ci.quantity,
			d.discount_rate,
			(ci.base_price * ci.quantity) AS normal_price,
			((ci.base_price - (ci.base_price * COALESCE(d.discount_rate, 0) / 100)) * ci.quantity) AS discount_price,
			pi.path AS image_path
		FROM 
			carts c
		JOIN 
			cart_items ci ON c.cart_id = ci.cart_id
		LEFT JOIN 
			product_images pi ON ci.product_id = pi.product_id
		JOIN 
			products p ON ci.product_id = p.id
		LEFT JOIN 
			discount d ON p.discount = d.discount_id
		WHERE 
			c.user_id = $1;
    `

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductCart])
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *ProductCartRepository) AddCart(ctx context.Context, req models.AddCartRequest) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, `
    WITH insert_cart AS (
        INSERT INTO carts (user_id)
        VALUES ($1)
        ON CONFLICT (user_id) DO UPDATE SET user_id = EXCLUDED.user_id 
        RETURNING cart_id
    ),
    current_cart AS (
        SELECT cart_id FROM insert_cart
        UNION ALL
        SELECT cart_id FROM carts WHERE user_id = $1
        LIMIT 1
    )
    INSERT INTO cart_items (cart_id, product_id, quantity, product_name, base_price, variant_name, size_name)
    SELECT cart_id, $2, $3, $4, $5, $6, $7 
    FROM current_cart
    ON CONFLICT (cart_id, product_id, variant_name, size_name)
    DO UPDATE SET
        quantity = cart_items.quantity + EXCLUDED.quantity;
    `,
		req.UserID,      // $1
		req.ProductID,   // $2
		req.Quantity,    // $3
		req.ProductName, // $4
		req.BasePrice,   // $5
		req.VariantName, // $6
		req.SizeName,    // $7
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
