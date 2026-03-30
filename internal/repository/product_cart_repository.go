package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductCartRepository struct {
	db *pgxpool.Pool
}

func NewProductCartRepository(db *pgxpool.Pool) *ProductCartRepository {
	return &ProductCartRepository{db: db}
}

func (r *ProductCartRepository) AddCart(ctx context.Context, userID int, productID int, quantity int, productName string, basePrice int, variantName string, sizeName string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, `
		WITH user_cart AS (
			INSERT INTO carts (user_id)
			VALUES ($7)
			ON CONFLICT (user_id)
			DO UPDATE SET user_id = EXCLUDED.user_id
			RETURNING cart_id
		)
		INSERT INTO cart_items (
			cart_id,
			product_id,
			quantity,
			product_name,
			base_price,
			variant_name,
			size_name
		)
		SELECT 
			cart_id,
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		FROM user_cart
		ON CONFLICT (cart_id, product_id, variant_name, size_name)
		DO UPDATE SET
			quantity = cart_items.quantity + EXCLUDED.quantity;
	`,
		productID,
		quantity,
		productName,
		basePrice,
		variantName,
		sizeName,
		userID,
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
