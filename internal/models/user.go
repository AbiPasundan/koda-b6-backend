package models

type Users struct {
	Id        int     `db:"id"`
	Full_Name string  `db:"full_name"`
	Email     string  `db:"email"`
	Address   *string `db:"address"`
	Phone     *string `db:"phone"`
}

// ("id", "full_name", "email", "password", "address", "phone", "pictures")
type User struct {
	Id        int    `db:"id"`
	Full_Name string `db:"full_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Address   string `db:"address"`
	Phone     string `db:"phone"`
	Pictures  string `db:"pictures"`
}

type Response struct {
	Success bool   `db:"success"`
	Message string `db:"message"`
	Results any    `db:"any"`
}
