package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type CategoryService struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepo: repo,
	}
}

func (p *CategoryService) GetCategory() ([]models.Category, error) {
	return p.CategoryRepo.GetCategory()
}

func (p *CategoryService) GetCategoryById(id int) (models.Category, error) {
	return p.CategoryRepo.GetCategoryById(id)
}

func (p *CategoryService) AddCategory(Category models.Category) (models.Category, error) {
	return p.CategoryRepo.AddCategory(Category)
}

func (p *CategoryService) UpdateCategoryById(id int, Category models.Category) (models.Category, error) {
	return p.CategoryRepo.UpdateCategoryById(id, Category)
}

func (p *CategoryService) DeleteCategoryById(id int) {
	p.CategoryRepo.DeleteCategoryById(id)
}
