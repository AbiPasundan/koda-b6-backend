package models

import "time"

type Users struct {
	Id        int     `json:"id" db:"id"`
	Full_Name string  `json:"full_name" db:"full_name"`
	Email     string  `json:"email" db:"email"`
	Address   *string `json:"address" db:"address"`
	Phone     *string `json:"phone" db:"phone"`
	Pictures  *string `json:"pictures" db:"pictures"`
}

// ("id", "full_name", "email", "password", "address", "phone", "pictures")
type User struct {
	Id        int     `db:"id" json:"id"`
	Full_Name string  `db:"full_name" json:"full_name"`
	Email     string  `db:"email" json:"email"`
	Password  string  `db:"password" json:"password"`
	Address   *string `db:"address" json:"address"`
	Phone     *string `db:"phone" json:"phone"`
	Pictures  *string `db:"pictures" json:"pictures"`
}

type ForgotPassword struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	// type after this cnage column name
	ExpiresAt time.Time `db:"expires_at"`
}

type AuthResetPassword struct {
	string `db:"expires_as"`
}

type Response struct {
	Success bool   `db:"success"`
	Message string `db:"message"`
	Results any    `db:"any"`
}

//	type UpdateProfile struct {
//		Id        int     `db:"id" json:"id"`
//		Full_Name string  `db:"full_name" json:"full_name"`
//		Email     string  `db:"email" json:"email"`
//		Password  string  `db:"password" json:"password"`
//		Address   *string `db:"address" json:"address"`
//		Phone     *string `db:"phone" json:"phone"`
//		Pictures  *string `db:"pictures" json:"pictures"`
//	}
type UpdateProfile struct {
	Id        int     `json:"id"`
	Full_Name *string `json:"full_name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	Address   *string `json:"address"`
	Phone     *string `json:"phone"`
	Pictures  *string `json:"pictures"`
}
