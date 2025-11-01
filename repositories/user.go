package repositories

import (
	"golang.org/x/crypto/bcrypt"
)

// UserRepository определяет интерфейс для работы с пользователями
type UserRepository interface {
	Register(user *models.User) error
	Authenticate(username, password string) (bool, error)
}

// InMemoryUserRepository — реализация репозитория с использованием map
type InMemoryUserRepository struct {
	users map[string]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]string),
	}
}

func (r *InMemoryUserRepository) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	r.users[user.Username] = string(hashedPassword)
	return nil
}

func (r *InMemoryUserRepository) Authenticate(username, password string) (bool, error) {
	storedPassword, exists := r.users[username]
	if !exists {
		return false, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
