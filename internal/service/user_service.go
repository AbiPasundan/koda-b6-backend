package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (s *UserService) GetUsers() ([]models.Users, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) GetUserById(id int) (models.User, error) {
	return s.UserRepo.GetUserById(id)
}

func (s *UserService) DeleteUserById(id int) {
	s.UserRepo.DeleteUserById(id)
}
