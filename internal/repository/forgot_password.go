package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type ForgotPasswordRepository struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository(db *pgx.Conn) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{db: db}
}

func (f *ForgotPasswordRepository) GetUserByEmail(token string) (models.ForgotPassword, error) {
	query := `SELECT email FROM users WHERE token = $1`
	rows, err := f.db.Query(context.Background(), query, token)
	if err != nil {
		return models.ForgotPassword{}, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
}

func (f *ForgotPasswordRepository) DeleteDataByEmail(db *pgx.Conn, code string) *ForgotPasswordRepository {
	query := `DELETE FROM forgot_password WHERE token = $1`
	_, err := f.db.Query(context.Background(), query, code)
	if err != nil {
		return nil
	}
	return &ForgotPasswordRepository{db: db}
}

func (f *ForgotPasswordRepository) CreateForgogtPasswordRequest(db *pgx.Conn, token string) *ForgotPasswordRepository {
	query := `INSERT INTO forgot_password (token) VALUES ($1)`
	_, err := f.db.Query(context.Background(), query, token)
	if err != nil {
		return nil
	}
	return &ForgotPasswordRepository{db: db}
}
