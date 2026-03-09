package models

import "time"

// -- id, delivery_method, full_name, email, address, sub_total, delivery_fee, tax, total, date, status, payment_method

var current_time string = time.Now().Format("2006-01-02 15:04:05")

type Order struct {
	Id            int32   `json:"id"`
	Delivery      string  `json:"delivery_method"`
	Name          string  `json:"full_name"`
	Email         string  `json:"email"`
	Address       string  `json:"address"`
	SubTotal      string  `json:"sub_total"`
	DeliveryFee   float32 `json:"delivery_fee"`
	Tax           float32 `json:"tax"`
	Total         float32 `json:"total"`
	Date          string  `json:"date"`
	Status        string  `json:"status"`
	PaymentMethod string  `json:"payment_method"`
}
