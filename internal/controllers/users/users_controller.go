package controllers

import (
	"encoding/json"
	"fmt"
	error_models "job_tracker/internal/models"
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

// Func to get all user information to display in account page
func GetUserInformation(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Get the user ID from the request query
	userId := r.URL.Query().Get("userId")

	// Validate and convert the user ID
	userIdInt, err := strconv.Atoi(userId)
	if err != nil || userIdInt <= 0 {
		// Invalid or missing user ID, respond with 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_models.ErrorResponse{Message: "Invalid or missing user ID"})
		return
	}

	// Call the service to get user information
	userInfo, err := services.GetUserInformation(userIdInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_models.ErrorResponse{Message: "Failed to retrieve user information"})
	}
	// if err != nil {
	// 	// Handle the case where the user information couldn't be retrieved
	// 	// If it's a "not found" error, return a 404 status code
	// 	if err == services.ErrUserNotFound { // Assume you have such an error defined
	// 		w.WriteHeader(http.StatusNotFound)
	// 		json.NewEncoder(w).Encode(error_models.ErrorResponse{Message: "User not found"})
	// 	} else {
	// 		// For other errors, return 500 Internal Server Error
	// 		w.WriteHeader(http.StatusInternalServerError)
	// json.NewEncoder(w).Encode(error_models.ErrorResponse{Message: "Failed to retrieve user information"})
	// 	}
	// 	return
	// }

	// Respond with the user information and 200 OK
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userInfo); err != nil {
		// If encoding fails, return 500 Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_models.ErrorResponse{Message: "Failed to encode response"})
	}
}

// Func UpdateUserPersonalDetails update personal information of user
func UpdateUserPersonalDetails(w http.ResponseWriter, r *http.Request) {
	var user_personal_info models.UserPersonalInfo

	// Decode the incoming user JSON
	err := json.NewDecoder(r.Body).Decode(&user_personal_info)
	fmt.Println(user_personal_info, "info")
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
	}

	fmt.Println(user_personal_info, "user udpate obj here")

	err = services.UpdateUserPersonalDetails(user_personal_info)

	if err != nil {
		fmt.Println(err, "err here for update")
		http.Error(w, "Failed to update user personal information", http.StatusInternalServerError)
		return
	}
	// Respond with the appropriate message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User personal information updated successfully",
	})
}

// Func UpdateUserPersonalDetails update personal information of user
func UpdateUserCareerDetails(w http.ResponseWriter, r *http.Request) {
	var user_career_info models.UserCareerInfo

	// Decode the incoming user JSON
	err := json.NewDecoder(r.Body).Decode(&user_career_info)
	fmt.Println(user_career_info, "info")
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
	}

	fmt.Println(user_career_info, "user udpate obj here")

	err = services.UpdateUserCareerDetails(user_career_info)

	if err != nil {
		fmt.Println(err, "err here for update")
		http.Error(w, "Failed to update user career information", http.StatusInternalServerError)
		return
	}
	// Respond with the appropriate message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User career information updated successfully",
	})
}
