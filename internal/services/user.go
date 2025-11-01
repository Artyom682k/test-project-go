package services

import (
	"github.com/Artyom682k/test-project-go/internal/models"
	"github.com/Artyom682k/test-project-go/internal/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) GetUserProfile(id int) (*models.UserProfile, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &models.UserProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
