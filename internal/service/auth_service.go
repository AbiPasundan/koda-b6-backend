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

func (p *AuthService) Register(user *models.AuthRegister) {
	p.AuthRepo.Register(user)
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
	userId, err := p.AuthRepo.GetUserIdByToken(req.Token)
	if err != nil {
		return err // Token invalid/expired
	}

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// if err != nil {
	//     return errors.New("gagal mengenkripsi password")
	// }

	err = p.AuthRepo.ResetPassword(userId, string(req.Password))
	if err != nil {
		return err
	}

	return nil
}
