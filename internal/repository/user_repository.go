package repository

import (
	"context"

	"backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAllUsers(conn *pgx.Conn) ([]models.Users, error) {

	rows, err := conn.Query(context.Background(), `
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
