package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	db *pgx.Conn
}

func NewAuthRepository(db *pgx.Conn) *AuthRepository {
	return &AuthRepository{db: db}
}

func (u *AuthRepository) FindEmail(email string) ([]models.AuthLogin, error) {

	rows, err := u.db.Query(context.Background(), `
		SELECT users.full_name, users.email
		FROM users
	`)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.AuthLogin])
}

func (u *AuthRepository) Register(user *models.AuthRegister) {
	_, err := u.db.Exec(context.Background(),
		"INSERT INTO users(full_name, email,password) VALUES($1,$2,$3) RETURNING (full_name, email", user.Email, user.Password, user.Full_Name,
	)
	if err != nil {
		return
	}

}
