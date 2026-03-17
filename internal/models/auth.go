package models

type AuthLogin struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
