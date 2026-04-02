package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"errors"
	"fmt"
)

type ProductCartService struct {
	ProductCartRepo *repository.ProductCartRepository
}

func NewProductCartService(repo *repository.ProductCartRepository) *ProductCartService {
	return &ProductCartService{
		ProductCartRepo: repo,
	}
}

func (s *ProductCartService) AddCart(ctx context.Context, req models.AddCartRequest) error {
	if req.UserID <= 0 {
		return errors.New("user ID tidak valid")
	}
	if req.ProductID <= 0 {
		return errors.New("product ID tidak valid")
	}
	if req.Quantity <= 0 {
		return errors.New("kuantitas harus lebih dari 0")
	}

	err := s.ProductCartRepo.AddCart(ctx, req)
	if err != nil {
		return fmt.Errorf("gagal menambahkan produk ke keranjang: %w", err)
	}

	return nil
}

// func (s *ProductCartService) GetCart(id int) (models.ProductCart, error) {
// 	return s.ProductCartRepo.GetCart(id)
// }

// Ubah return type menjadi []models.ProductCart
func (s *ProductCartService) GetCart(id int) ([]models.ProductCart, error) {
	return s.ProductCartRepo.GetCart(id)
}
