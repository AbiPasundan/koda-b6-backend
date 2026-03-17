package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"fmt"

	"github.com/matthewhartstonge/argon2"
)

type AuthService struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepo: repo,
	}
}

func (p *CategoryService) FindEmail(email string, password []byte) ([]models.AuthLogin, error) {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return nil, err
	}
	fmt.Println(string(encoded))

	return p.CategoryRepo.FindEmail(email)
}
