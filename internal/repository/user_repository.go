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

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetUserById(id int) (models.User, error) {
	rows, err := u.db.Query(context.Background(), `
		SELECT id, full_name, email, password, address, phone, pictures FROM users WHERE id = $1
	`, id)

	if err != nil {
		return models.User{}, err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	return result, err
}
