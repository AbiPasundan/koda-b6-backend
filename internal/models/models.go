package models

type Users struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type User struct {
	Id        int    `json:"id"`
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}
