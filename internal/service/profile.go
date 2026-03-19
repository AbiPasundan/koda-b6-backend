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

func (s *ProfileService) GetProfile(data models.User) (*models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// func (s *ProfileService) UpdateProfile(data *models.User) (*models.User, error) {
// 	return s.UserRepo.UpdateUserById(data.Id, *data)
// }
