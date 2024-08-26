package controllers

import (
	"encoding/json"
	models "job_tracker/internal/models/users"
	services "job_tracker/internal/services/users"
	validation "job_tracker/internal/validation/users"
	"net/http"
)

// CreateUser handles the user creation endpoint
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the incoming user JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !validation.IsValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Validate name length (should be longer than 2 characters)
	if len(user.Name) <= 2 {
		http.Error(w, "Name must be longer than 2 characters", http.StatusBadRequest)
		return
	}

	// Create the user
	createdUser, isNewUser, err := services.CreateNewUser(user.Email, user.Name)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Respond with the appropriate message
	if isNewUser {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "User created successfully",
			"user":     createdUser,
			"new_user": true,
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "User already exists",
			"user":     createdUser,
			"new_user": false,
		})
	}
}
