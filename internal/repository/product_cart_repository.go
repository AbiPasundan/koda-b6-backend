package repository

import (
	"backend/internal/models"
	"context"
	"errors"

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
		SELECT  ci.cart_item_id, ci.product_id, ci.product_name, ci.variant_name, ci.size_name, ci.base_price, ci.quantity, d.discount_rate, d.is_flash_sale, (ci.base_price * ci.quantity) AS normal_price, ((ci.base_price - (ci.base_price * COALESCE(d.discount_rate, 0) / 100)) * ci.quantity) AS discount_price, pi.path AS image_path
		FROM carts c
		JOIN cart_items ci ON c.cart_id = ci.cart_id
		JOIN products p ON ci.product_id = p.id
		LEFT JOIN discount d ON p.discount = d.discount_id

		LEFT JOIN LATERAL ( SELECT path FROM product_images WHERE product_id = ci.product_id ORDER BY id ASC LIMIT 1 ) pi ON true

		WHERE c.user_id = $1;
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
func (r *ProductCartRepository) DeleteCart(id int) error {
	query := `DELETE FROM cart_items WHERE cart_item_id = $1`

	result, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("no cart found with this id")
	}

	return nil
}

func (r *ProductCartRepository) GetOrder(id int) ([]models.HistoryOrder, error) {
	query := `select orders.user_id, orders.status, orders.total, orders.image_path, orders.created_at from orders where user_id = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.HistoryOrder])
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *ProductCartRepository) AddOrder(ctx context.Context, userID int) (int, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return userID, err
	}
	defer tx.Rollback(ctx)

	var orderID string
	err = tx.QueryRow(ctx, `
		INSERT INTO orders (user_id, status, total, image_path)
		SELECT 
			c.user_id, 'pending', SUM(ci.base_price * ci.quantity),
			(SELECT pi.path FROM cart_items ci2 
			 LEFT JOIN LATERAL (SELECT path FROM product_images WHERE product_id = ci2.product_id LIMIT 1) pi ON TRUE 
			 WHERE ci2.cart_id = c.cart_id LIMIT 1)
		FROM carts c
		JOIN cart_items ci ON ci.cart_id = c.cart_id
		WHERE c.user_id = $1
		GROUP BY c.user_id, c.cart_id
		RETURNING id
	`, userID).Scan(&orderID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return userID, err
		}
		return userID, err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO order_items (order_id, product_id, product_name, quantity, price, size, variant, image_path)
		SELECT $1, ci.product_id, ci.product_name, ci.quantity, ci.base_price, ci.size_name, ci.variant_name, pi.path
		FROM cart_items ci
		JOIN carts c ON ci.cart_id = c.cart_id
		LEFT JOIN LATERAL (SELECT path FROM product_images WHERE product_id = ci.product_id LIMIT 1) pi ON TRUE
		WHERE c.user_id = $2
	`, orderID, userID)
	if err != nil {
		return userID, err
	}

	_, err = tx.Exec(ctx, `
		DELETE FROM cart_items
		USING carts
		WHERE cart_items.cart_id = carts.cart_id
		AND carts.user_id = $1
	`, userID)
	if err != nil {
		return userID, err
	}

	return userID, tx.Commit(ctx)
}
