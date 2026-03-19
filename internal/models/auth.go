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
	Email string `json:"email" binding:"required,email"`
}

// type AuthResetPassword struct {
// 	Id       int    `json:"id"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type ForgotPasswordRequest struct {
// 	Email string `json:"email" binding:"required,email"`
// }

type ResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
