package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *AuthRepository {
	return &AuthRepository{db: db}
}

func (u *CategoryRepository) FindEmail() ([]models.AuthLogin, error) {

	rows, err := u.db.Query(context.Background(), `
		SELECT users.full_name, users.email
		FROM users
	`)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.AuthLogin])
}
