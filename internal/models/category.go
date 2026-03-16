package models

type Category struct {
	Id   int    `json:"category_id" db:"category_id"`
	Name string `json:"category_name" db:"category_name"`
}
