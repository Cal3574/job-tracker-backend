package controllers

import (
	"encoding/json"
	"fmt"
	models "job_tracker/internal/models/users"
	services "job_tracker/internal/services/users"
	validation "job_tracker/internal/validation/users"
	"net/http"
	"strconv"
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

// Func to complete the users sign up process
func CompleteSignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the incoming user JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	//check the value of the userId

	if user.ID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Call the service to complete the user sign up
	err = services.CompleteSignUp(user)
	if err != nil {
		fmt.Println(err, "ERROR IS HERE")
		http.Error(w, "Failed to complete sign up", http.StatusInternalServerError)
		return
	}

	// Respond with the appropriate message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Sign up completed successfully",
	})

}

// Func to check if the user has completed the sign up process

func CheckUserSignUpStatus(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request context
	userId := r.URL.Query().Get("userId")

	// Turn the user ID into an integer
	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the service to check the sign up status
	signUpStatus, err := services.CheckUserSignUpStatus(userIdInt)
	if err != nil {
		http.Error(w, "Failed to check sign up status", http.StatusInternalServerError)
		return
	}

	// Respond with the appropriate message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"sign_up_complete": signUpStatus,
	})
}
