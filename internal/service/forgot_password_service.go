package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"math/rand"
	"strconv"
)

// * RequestForgotPassword
//   * pastikan email ada dari repo user
//   * generate code
//   * insert (ke repo forgot_password)
// * reset password
//   * pastikan code dan email ada
//   * update pass (dari repo user)
//   * delete dari repo forgot_password

type ForgotPasswordService struct {
	ForgotPasswordRepo *repository.ForgotPasswordRepository
	UserRepo           *repository.UserRepository
}

func NewForgotPasswordService(fpRepo *repository.ForgotPasswordRepository, userRepo *repository.UserRepository) *ForgotPasswordService {
	return &ForgotPasswordService{
		ForgotPasswordRepo: fpRepo,
		UserRepo:           userRepo,
	}
}

func (f *ForgotPasswordService) RequestForgotPassword(req models.JustEmail) (string, error) {
	user, err := f.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	code := strconv.Itoa(rand.Intn(999999))

	_ = f.ForgotPasswordRepo.DeleteCode(user.Id)
	err = f.ForgotPasswordRepo.CreateForgotPassword(user.Id, code)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (f *ForgotPasswordService) ResetPassword(req models.ResetPasswordInput) error {
	user, err := f.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	_, err = f.ForgotPasswordRepo.GetTokenByUserIdAndCode(user.Id, req.Code)
	if err != nil {
		return err
	}

	err = f.UserRepo.UpdatePasswordByEmail(req.Email, req.NewPassword)
	if err != nil {
		return err
	}

	return f.ForgotPasswordRepo.DeleteCode(user.Id)
}
