package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{
		OrderRepo: repo,
	}
}
func (p *OrderService) GetOrder() ([]models.Order, error) {
	return p.OrderRepo.GetAllOrder()
}
