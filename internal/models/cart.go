package models

// cart_id, user_id, quantity, size_name, total_price, items[]

type Cart struct {
	CartId   int `json:"cart_id" db:"cart_id"`
	UserId   int `json:"user_id" db:"user_id"`
	Quantity int `json:"quantity" db:"quantity"`
	CartItem CartItem
	// Products   []Item `json:"items" db:"items"`
}
