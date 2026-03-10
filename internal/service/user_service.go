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

func (u *UserService) GetUsers() ([]models.Users, error) {
	return u.UserRepo.GetAllUsers()
}

func (u *UserService) GetUserById(id int) (models.User, error) {
	return u.UserRepo.GetUserById(id)
}

func (u *UserService) AddUser(user models.User) (models.User, error) {
	return u.UserRepo.AddUser(user)
}

func (u *UserService) DeleteUserById(id int) {
	u.UserRepo.DeleteUserById(id)
}
