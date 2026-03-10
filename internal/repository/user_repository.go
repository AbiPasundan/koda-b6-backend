package repository

import (
	"context"

	"backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetAllUsers() ([]models.Users, error) {

	rows, err := u.db.Query(context.Background(), `
		SELECT id, full_name, email, address, phone
		FROM users
	`)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])
}

func (u *UserRepository) GetUserById(id int) (models.User, error) {
	rows, err := u.db.Query(context.Background(), `
		SELECT id, full_name, email, password, address, phone, pictures FROM users WHERE id = $1
	`, id)

	if err != nil {
		return models.User{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
}

func (u *UserRepository) AddUser(user models.User) (models.User, error) {
	query := `
		INSERT INTO users (full_name, email, password, address, phone, pictures)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, full_name, email, password, address, phone, pictures
	`
	rows, err := u.db.Query(context.Background(), query, user.Full_Name, user.Email, user.Password, user.Address, user.Phone, user.Pictures)

	if err != nil {
		return models.User{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
}

// update
func (u *UserRepository) UpdateUserById(id int, user models.User) (models.User, error) {
	// note jika update harus kirim model untuk update data
	rows, err := u.db.Query(context.Background(), `
		UPDATE users SET full_name = $1, email = $2, password = $3, address = $4, phone = $5, pictures = $6 WHERE id = $7 RETURNING id, full_name, email, password, address, phone, pictures
	`, user.Full_Name, user.Email, user.Password, user.Address, user.Phone, user.Pictures, id)

	// rows, err := u.db.Query(context.Background(), `
	// 	UPDATE users SET full_name = $1, email = $2, password = $3, address = $4, phone = $5, pictures = $6 WHERE id = $7 RETURNING id, full_name, email, password, address, phone, pictures
	// `, user.Full_Name, user.Email, user.Password, user.Address, user.Phone, user.Pictures, id)

	if err != nil {
		return models.User{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
}

func (u *UserRepository) DeleteUserById(id int) {
	rows, err := u.db.Query(context.Background(), `
		DELETE FROM users WHERE id = $1;
	`, id)

	if err != nil {
		return
	}

	defer rows.Close()
}
