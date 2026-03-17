package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type AuthService struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepo: repo,
	}
}

func (p *AuthService) FindEmail(email string) ([]models.AuthLogin, error) {
	return p.AuthRepo.FindEmail(email)
}

func (p *AuthService) Register(email *models.AuthRegister) {
	p.AuthRepo.Register(email)
}
