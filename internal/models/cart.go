package models

// cart_id, user_id, quantity, size_name, total_price, items[]

type Cart struct {
	CartId   int `json:"cart_id" db:"cart_id"`
	UserId   int `json:"user_id" db:"user_id"`
	Quantity int `json:"quantity" db:"quantity"`
	CartItem CartItem
	// Products   []Item `json:"items" db:"items"`
}

type AddCartRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	ProductID   int    `json:"product_id" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	ProductName string `json:"product_name"`
	BasePrice   int    `json:"base_price"`
	VariantName string `json:"variant_name"`
	SizeName    string `json:"size_name"`
}
