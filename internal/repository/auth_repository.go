package repository

import (
	"backend/internal/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (u *AuthRepository) FindByEmail(email string) (*models.AuthLogin, error) {
	row := u.db.QueryRow(context.Background(), `
		SELECT id, email, password 
		FROM users 
		WHERE email = $1;
	`, email)

	var user models.AuthLogin
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *AuthRepository) Register(user *models.AuthRegister) {
	_, err := u.db.Exec(context.Background(),
		"INSERT INTO users(full_name, email, password) VALUES($1, $2, $3)", user.Full_Name, user.Email, user.Password,
	)
	if err != nil {
		log.Println(err)
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
	query := `INSERT INTO forgot_password (user_id, token, created_at, expires_at) VALUES ($1, $2, NOW(), NOW() + INTERVAL '5 minutes')`
	_, err := f.db.Exec(context.Background(), query, userId, token)
	return err
}

func (f *AuthRepository) GetByToken(token string) (models.ForgotPassword, error) {
	var data models.ForgotPassword

	query := `
	SELECT user_id, expires_at 
	FROM forgot_password 
	WHERE token = $1
	`

	err := f.db.QueryRow(context.Background(), query, token).
		Scan(&data.UserId, &data.ExpiresAt)

	return data, err
}

func (f *AuthRepository) ResetPassword(userId int, password string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := f.db.Exec(context.Background(), query, password, userId)
	return err
}

func (f *AuthRepository) DeleteToken(token string) error {
	query := `DELETE FROM forgot_password WHERE token = $1`
	_, err := f.db.Exec(context.Background(), query, token)
	return err
}
