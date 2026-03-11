package service

import (
	"backend/internal/repository"
	"strconv"

	"math/rand/v2"
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

func (f *ForgotPasswordService) RequestForgotPassword(email string) (string, error) {
	if _, err := f.ForgotPasswordRepo.GetUserByEmail(email); err != nil {
		return "", err
	}
	code := strconv.Itoa(rand.IntN(999999))
	f.ForgotPasswordRepo.CreateForgogtPasswordRequest(code)
	return code, nil
}

func (f *ForgotPasswordService) ResetPassword(email string, code string, newPassword string) error {
	_, err := f.ForgotPasswordRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	f.UserRepo.UpdatePasswordByEmail(email, newPassword)
	f.ForgotPasswordRepo.DeleteDataByEmail(email)
	return err
}
