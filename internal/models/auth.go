package models

type AuthLogin struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRegister struct {
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AuthForgotPassword struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
