package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductCartRepository struct {
	db *pgxpool.Pool
}

func NewProductCartRepository(db *pgxpool.Pool) *ProductCartRepository {
	return &ProductCartRepository{db: db}
}

type CartItem struct {
	CartID      int              `json:"cart_id"`
	ProductID   int              `json:"product_id"`
	Quantity    int              `json:"quantity"`
	ProductName string           `json:"product_name"`
	BasePrice   int              `json:"base_price"`
	VariantName []models.Variant `json:"variant_name"`
	SizeName    []models.Size    `json:"size_name"`
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
