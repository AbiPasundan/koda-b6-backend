package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"fmt"
	"math/rand"
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

func (p *AuthService) ForgotPasswordRequest(req *models.AuthForgotPassword) error {

	user, err := p.AuthRepo.GetEmail(req.Email)
	if err != nil {
		return err
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	fmt.Printf("OTP Request %s: %s\n", req.Email, code)

	return p.AuthRepo.RequestForgotPassword(user.Id, code)
}
