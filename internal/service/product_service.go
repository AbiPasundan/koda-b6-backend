package service

import (
	"backend/internal/models"
	"backend/internal/repository"

	"github.com/jackc/pgx/v5"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepo: repo,
	}
}

func (p *ProductService) GetProduct(conn *pgx.Conn) ([]models.Product, error) {
	return p.ProductRepo.GetAllProduct(conn)
}
