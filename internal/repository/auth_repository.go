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
		SELECT email, password
		FROM users
		WHERE email = $1
	`, email)
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

func (f *AuthRepository) GetEmail(email string) (*models.User, error) {
	query := `SELECT id, full_name, email, password, address, phone, pictures FROM users WHERE email = $1`
	rows, err := f.db.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (f *AuthRepository) RequestForgotPassword(userId int, token string) error {
	query := `INSERT INTO forgot_password (user_id, token, created_at) VALUES ($1, $2, NOW())`
	_, err := f.db.Exec(context.Background(), query, userId, token)
	return err
}

func (f *AuthRepository) ResetPassword(userId int, token string) error {
	query := `INSERT INTO forgot_password (user_id, token, created_at) VALUES ($1, $2, NOW())`
	_, err := f.db.Query(context.Background(), query, userId, token)
	return err
}
