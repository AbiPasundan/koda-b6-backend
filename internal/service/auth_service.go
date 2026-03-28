package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type AuthService struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepo: repo,
	}
}

func (s *AuthService) Login(email, password string) (*models.AuthLogin, error) {
	user, err := s.AuthRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) Register(user *models.AuthRegister) error {
	user.Password = string(user.Password)

	return s.AuthRepo.Register(user)
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

func (p *AuthService) ResetPassword(req models.ResetPasswordRequest) error {

	data, err := p.AuthRepo.GetByToken(req.Token)
	if err != nil {
		return errors.New("invalid token")
	}

	if time.Now().After(data.ExpiresAt) {
		return errors.New("token expired")
	}

	err = p.AuthRepo.ResetPassword(data.UserId, string(req.Password))
	if err != nil {
		return err
	}

	return p.AuthRepo.DeleteToken(req.Token)
}
