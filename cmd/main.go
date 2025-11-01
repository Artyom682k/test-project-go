package main

import (
	"fmt"
	"net/http"

	"github.com/Artyom682k/test-project-go/handlers"
	"github.com/Artyom682k/test-project-go/repositories"
	"github.com/Artyom682k/test-project-go/services"
)

func main() {
	// Инициализация репозитория
	userRepo := repositories.NewInMemoryUserRepository()

	// Инициализация сервиса
	userService := services.NewUserService(userRepo)

	// Регистрация обработчиков
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, userService)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, userService)
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
