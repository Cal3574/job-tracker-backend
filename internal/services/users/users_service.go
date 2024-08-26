package users

import (
	models "job_tracker/internal/models/users"
	repositories "job_tracker/internal/repositories/users"
)

// CreateNewUser creates a new user record or returns an existing user.
func CreateNewUser(email string, name string) (models.User, bool, error) {
	// Create the user model
	user := models.User{
		Email: email,
		Name:  name,
	}

	// Call the repository layer to create the user or get the existing one
	createdUser, isNewUser, err := repositories.CreateUser(user)
	if err != nil {
		// Pass the error up to the controller
		return models.User{}, false, err
	}

	return createdUser, isNewUser, nil
}
