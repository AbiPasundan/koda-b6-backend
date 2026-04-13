package models

import "time"

// -- id, delivery_method, full_name, email, address, sub_total, delivery_fee, tax, total, date, status, payment_method

var current_time string = time.Now().Format("2006-01-02 15:04:05")

// SELECT orders.id, orders.user_id, orders.status, orders.total, orders.image_path, orders.updated_at FROM orders;
type Order struct {
	Id        string    `json:"id" db:"id"`
	UserId    int       `json:"user_id db:"user_id`
	Status    string    `json:"status" db:"status"`
	Total     int       `json:"total" db:"total"`
	ImagePath string    `json:"image_path" db:"image_path"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
