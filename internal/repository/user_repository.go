package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewUserRepository(db *pgxpool.Pool, rdb *redis.Client) *UserRepository {
	return &UserRepository{db: db, rdb: rdb}
}

func (u *UserRepository) GetAllUsers() ([]models.Users, error) {
	ctx := context.Background()
	cacheKey := "users:all"

	val, err := u.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var users []models.Users
		if err := json.Unmarshal([]byte(val), &users); err == nil {
			log.Println("Cache HIT: users")
			return users, nil
		}
	}

	log.Println("Cache MISS: users")

	rows, err := u.db.Query(ctx, `
		SELECT id, full_name, email, address, phone, pictures FROM users;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return nil, err
	}

	// data, err := json.Marshal(users)
	// if err == nil {
	// 	u.rdb.Set(ctx, cacheKey, data, 5*time.Minute)
	// }
	if err := json.Unmarshal([]byte(val), &users); err == nil {
		log.Println("Cache HIT: users")
		return users, nil
	} else {
		log.Println("Unmarshal error:", err)
	}

	return users, nil
}

// func (u *UserRepository) GetAllUsers() ([]models.Users, error) {

// 	rows, err := u.db.Query(context.Background(), `
// 		SELECT id, full_name, email, address, phone, pictures FROM users;
// 	`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])
// }

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

func (u *UserRepository) UpdateUserById(id int, user models.User) (models.User, error) {
	rows, err := u.db.Query(context.Background(), `
		UPDATE users SET full_name = $1, email = $2, password = $3, address = $4, phone = $5, pictures = $6 WHERE id = $7 RETURNING id, full_name, email, password, address, phone, pictures
	`, user.Full_Name, user.Email, user.Password, user.Address, user.Phone, user.Pictures, id)

	if err != nil {
		return models.User{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
}

func (u *UserRepository) DeleteUserById(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := u.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("no user found with this id")
	}

	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {
	rows, err := u.db.Query(context.Background(), `
		SELECT id, full_name, email, password, address, phone, pictures FROM users WHERE email = $1
	`, email)

	if err != nil {
		return models.User{}, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
}

func (u *UserRepository) UpdatePasswordByEmail(email string, newPassword string) error {
	_, err := u.db.Query(context.Background(), `
		UPDATE users SET password = $1 WHERE email = $2
	`, newPassword, email)

	return err
}

// func (u *UserRepository) UpdateProfile(id int, user models.UpdateProfile) (models.UpdateProfile, error) {
// 	query := `UPDATE users SET full_name = COALESCE($1, full_name), email = COALESCE($2, email), password = COALESCE($3, password), address = COALESCE($4, address), phone = COALESCE($5, phone), pictures = COALESCE($6, pictures) WHERE id = $7 RETURNING id, full_name, email, password, address, phone, pictures`
// 	// rows, err := u.db.Query(context.Background(), `
// 	// 	UPDATE users SET full_name = $1, email = $2, password = $3, address = $4, phone = $5, pictures = $6 WHERE id = $7 RETURNING id, full_name, email, password, address, phone, pictures
// 	// `, user.Full_Name, user.Email, user.Password, user.Address, user.Phone, user.Pictures, id)
// 	rows, err := u.db.Query(context.Background(), query, user.Full_Name, user.Email, user.Password, user.Address, user.Phone, user.Pictures, id)

// 	if err != nil {
// 		return models.UpdateProfile{}, err
// 	}

// 	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UpdateProfile])
// }

func (u *UserRepository) UpdateProfile(id int, user models.UpdateProfile) (models.UpdateProfile, error) {
	query := `
	UPDATE users SET
		full_name = COALESCE($1, full_name),
		email = COALESCE($2, email),
		password = COALESCE($3, password),
		address = COALESCE($4, address),
		phone = COALESCE($5, phone),
		pictures = COALESCE($6, pictures)
	WHERE id = $7
	RETURNING id, full_name, email, password, address, phone, pictures
	`

	row := u.db.QueryRow(context.Background(),
		query,
		user.Full_Name,
		user.Email,
		user.Password,
		user.Address,
		user.Phone,
		user.Pictures,
		id,
	)

	var result models.UpdateProfile
	err := row.Scan(
		&result.Id,
		&result.Full_Name,
		&result.Email,
		&result.Password,
		&result.Address,
		&result.Phone,
		&result.Pictures,
	)

	if err != nil {
		return models.UpdateProfile{}, err
	}

	return result, nil
}
