package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthLogin struct {
	Id        int        `json:"id"`
	Email     string     `json:"email"`
	Full_Name string     `json:"full_name"`
	Password  string     `json:"password"`
	Address   *string    `db:"address" json:"address"`
	Phone     *string    `db:"phone" json:"phone"`
	Pictures  *string    `db:"pictures" json:"pictures"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	Role      string     `json:"role"`
}

type AuthRegister struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthForgotPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type Claims struct {
	UserID    int        `db:"user_id" json:"user_id"`
	Email     string     `db:"email" json:"email"`
	FullName  string     `db:"full_name" json:"full_name"`
	Address   *string    `db:"address" json:"address"`
	Phone     *string    `db:"phone" json:"phone"`
	Pictures  *string    `db:"pictures" json:"pictures"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	Role      string     `db:"role" json:"role"`
	jwt.RegisteredClaims
}
