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

	// code, err := f.generateCode()
	// if err != nil {
	// 	return "", err
	// }

	// rawCode := rand.IntN(999999)
	// test := strconv.Itoa(rawCode)

	code := strconv.Itoa(rand.IntN(999999))

	f.ForgotPasswordRepo.CreateForgogtPasswordRequest(code)

	return code, nil
}

func (f *ForgotPasswordService) ResetPassword(email string, code string, newPassword string) error {
	_, err := f.ForgotPasswordRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	// email, err := f.ForgotPasswordRepo.GetUserByEmail(email)
	// if email == "test@mail.com" {
	// 	return err
	// } else if err != nil {
	// 	return err
	// }
	// test, err := f.UserRepo.GetUserById(1)

	// err := f.ForgotPasswordRepo.VerifyCode(email, code)
	// if err != nil {
	// 	return err
	// }

	// err = f.UserRepo.UpdatePassword(email, newPassword)
	// if err != nil {
	// 	return err
	// }

	return err

}
