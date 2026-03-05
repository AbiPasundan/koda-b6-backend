package service

import (
	"backend/internal/models"
	"backend/internal/repository"

	"github.com/jackc/pgx/v5"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (s *UserService) GetUsers(conn *pgx.Conn) ([]models.Users, error) {
	return s.UserRepo.GetAllUsers(conn)
}
