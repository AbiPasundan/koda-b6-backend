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

// cart_id, product_id, quantity, product_name, base_price, variant_name, size_name

type CartItem struct {
	CartID      int              `json:"cart_id"`
	ProductID   int              `json:"product_id"`
	Quantity    int              `json:"quantity"`
	ProductName string           `json:"product_name"`
	BasePrice   int              `json:"base_price"`
	VariantName []models.Variant `json:"variant_name"`
	SizeName    []models.Size    `json:"size_name"`
}

func (r *ProductCartRepository) AddCart(ctx context.Context, userID int, productId int, productName string, productPrice int, quantity int, variantName string, SizeName string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, `
	INSERT INTO carts (user_id)
	VALUES ($1)
	ON CONFLICT (user_id) DO NOTHING;

	WITH user_cart AS (SELECT cart_id FROM carts WHERE user_id = $1)

	INSERT INTO cart_items ( cart_id, product_id, quantity, product_name, base_price, variant_name, size_name)
	SELECT  cart_id, $2, $3, $4, $5, $6, $7
	FROM user_cart

	ON CONFLICT (cart_id, product_id, variant_name, size_name)
	DO UPDATE SET
		quantity = cart_items.quantity + EXCLUDED.quantity;
	`,
		userID,
		productId,
		quantity,
		productName,
		productPrice,
		variantName,
		SizeName, // BasePrice string, VariantName string, SizeName string
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
