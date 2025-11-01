package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "strings"

	"github.com/Artyom682k/test-project-go/models"
	"github.com/Artyom682k/test-project-go/services"
)

// RegisterHandler обрабатывает регистрацию
func RegisterHandler(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := userService.Register(&user); err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s registered successfully\n", user.Username)
}

// LoginHandler обрабатывает авторизацию
func LoginHandler(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	authenticated, err := userService.Authenticate(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	if !authenticated {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "User %s authenticated successfully\n", user.Username)
}
