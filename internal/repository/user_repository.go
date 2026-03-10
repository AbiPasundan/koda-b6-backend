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

func (u *UserRepository) GetAllUsers(db *pgx.Conn) ([]models.Users, error) {

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
