package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type ForgotPasswordRepository struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository(db *pgx.Conn) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{db: db}
}

func (f *ForgotPasswordRepository) CreateForgotPassword(userId int, token string) error {
	query := `INSERT INTO forgot_password (user_id, token, created_at) VALUES ($1, $2, NOW())`
	_, err := f.db.Query(context.Background(), query, userId, token)
	return err
}

func (f *ForgotPasswordRepository) GetTokenByUserIdAndCode(userId int, code string) (string, error) {
	var token string
	query := `SELECT token FROM forgot_password WHERE user_id = $1 AND token = $2`
	err := f.db.QueryRow(context.Background(), query, userId, code).Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (f *ForgotPasswordRepository) DeleteCode(userId int) error {
	query := `DELETE FROM forgot_password WHERE user_id = $1`
	_, err := f.db.Exec(context.Background(), query, userId)
	return err
}
