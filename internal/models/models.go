package models

type Users struct {
	Id        int    `json:"id"`
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

// ("id", "full_name", "email", "password", "address", "phone", "pictures")
type User struct {
	Id        int    `json:"id"`
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Pictures  string `json:"pictures"`
}
