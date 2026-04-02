package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"errors"
	"fmt"
	"math/rand"
	"time"

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

func (s *AuthService) Login(email, password string) (*models.AuthLogin, error) {
	// argon := argon2.DefaultConfig()

	user, err := s.AuthRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	match, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("error verifying password: %w", err)
	}

	if !match {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) Register(user *models.AuthRegister) error {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(user.Password))

	if err != nil {
		panic(err)
	}

	user.Password = string(encoded)

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
