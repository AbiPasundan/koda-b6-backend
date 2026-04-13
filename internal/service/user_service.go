package service

import (
	"backend/internal/models"
	"backend/internal/repository"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	UserRepo *repository.UserRepository
	rdb      *redis.Client
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

func (u *UserService) DeleteUserById(id int) error {
	return u.UserRepo.DeleteUserById(id)
}

func (u *UserService) UpdateUserById(id int, user models.User) {
	u.UserRepo.UpdateUserById(id, user)
}
