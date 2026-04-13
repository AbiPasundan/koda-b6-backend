package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type ProfileService struct {
	UserRepo *repository.UserRepository
}

func NewProfileService(repo *repository.UserRepository) *ProfileService {
	return &ProfileService{
		UserRepo: repo,
	}
}

func (s *ProfileService) GetProfile(email string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}

func (u *UserService) UpdateUser(id int, user models.User) {
	u.UserRepo.UpdateProfile(id, user)
}
