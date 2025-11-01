package services

import (
	"errors"
	"github.com/Artyom682k/test-project-go/models"
	"github.com/Artyom682k/test-project-go/repositories"
)

// UserService содержит логику бизнес-операций
type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(user *models.User) error {
	return s.userRepo.Register(user)
}

func (s *UserService) Authenticate(username, password string) (bool, error) {
	return s.userRepo.Authenticate(username, password)
}
