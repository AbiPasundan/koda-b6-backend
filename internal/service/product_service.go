package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepo: repo,
	}
}

func (p *ProductService) GetProduct() ([]models.Product, error) {
	return p.ProductRepo.GetAllProduct()
}

func (p *ProductService) GetProductById(id int) (models.Product, error) {
	return p.ProductRepo.GetProductById(id)
}

func (p *ProductService) AddProduct(product models.Product) (models.Product, error) {
	return p.ProductRepo.AddProduct(product)
}

func (p *ProductService) UpdateProductById(id int, product models.Product) (models.Product, error) {
	return p.ProductRepo.UpdateProductById(id, product)
}

func (p *ProductService) DeleteProductById(id int) error {
	return p.ProductRepo.DeleteProductById(id)
}
