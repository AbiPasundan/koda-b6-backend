package repository

import (
	"context"

	"backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type CategoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (u *CategoryRepository) GetCategory() ([]models.Category, error) {

	rows, err := u.db.Query(context.Background(), `
		SELECT category_id, category_name
		FROM category
	`)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Category])
}

func (u *CategoryRepository) GetCategoryById(id int) (models.Category, error) {
	rows, err := u.db.Query(context.Background(), `
		SELECT category_id, category_name FROM category WHERE category_id = $1
	`, id)

	if err != nil {
		return models.Category{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
}

func (u *CategoryRepository) AddCategory(category models.Category) (models.Category, error) {
	query := `
		INSERT INTO category (category_name)
		VALUES ($1)
		RETURNING category_id, category_name
	`
	rows, err := u.db.Query(context.Background(), query, category.Name)

	if err != nil {
		return models.Category{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
}

func (u *CategoryRepository) UpdateCategoryById(id int, category models.Category) (models.Category, error) {
	rows, err := u.db.Query(context.Background(), `
		UPDATE category SET category_name = $1 WHERE category_id = $2 RETURNING category_id, category_name 
	`, category.Name, id)

	if err != nil {
		return models.Category{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
}

func (u *CategoryRepository) DeleteCategoryById(id int) {
	rows, err := u.db.Query(context.Background(), `
		DELETE FROM category WHERE category_id = $1;
	`, id)

	if err != nil {
		return
	}

	defer rows.Close()
}
