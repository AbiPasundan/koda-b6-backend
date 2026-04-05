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
func (s *ProductCartService) GetCart(id int) ([]models.ProductCart, error) {
	return s.ProductCartRepo.GetCart(id)
}

func (s *ProductCartService) DeleteCartById(id int) error {
	return s.ProductCartRepo.DeleteCart(id)
}

func (s *ProductCartService) GetOrder(id int) ([]models.HistoryOrder, error) {
	return s.ProductCartRepo.GetOrder(id)
}

func (s *ProductCartService) AddOrder(ctx context.Context, userID int) (int, error) {

	if userID == 0 {
		return userID, fmt.Errorf("Login Heula")
	}

	orderID, err := s.ProductCartRepo.AddOrder(ctx, userID)
	if err != nil {
		return userID, err
	}

	return orderID, nil
}

func (s *ProductCartService) GetOrderById(id string) ([]models.DetailOrder, error) {
	return s.ProductCartRepo.GetOrderById(id)
}
