package service

import (
	"backend/internal/models"
	"backend/internal/repository"

	"github.com/matthewhartstonge/argon2"
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
func (u *ProfileService) UpdateUser(id int, user models.UpdateProfile) (models.UpdateProfile, error) {

	if user.Password != nil {

		argon := argon2.DefaultConfig()

		encoded, err := argon.HashEncoded([]byte(*user.Password))
		if err != nil {
			return models.UpdateProfile{}, err
		}

		hash := string(encoded)
		user.Password = &hash
	}

	return u.UserRepo.UpdateProfile(id, user)
}
