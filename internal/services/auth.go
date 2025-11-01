package services

import (
	"github.com/Artyom682k/test-project-go/internal/models"
	"github.com/Artyom682k/test-project-go/internal/repositories"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// Здесь должна быть логика проверки пароля
	// Например: if !checkPassword(user.PasswordHash, password) { ... }

	return user, nil
}
