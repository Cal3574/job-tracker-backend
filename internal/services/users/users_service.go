package users

import (
	"fmt"
	models "job_tracker/internal/models/users"
	repositories "job_tracker/internal/repositories/users"
)

// CreateNewUser creates a new user record by delegating to the repository layer.
func CreateNewUser(email string, name string) (models.User, error) {
	fmt.Println("Creating new user")

	// Create the user model
	user := models.User{
		Email: email,
		Name:  name,
	}

	// Call the repository layer to create the user
	createdUser, err := repositories.CreateUser(user)
	if err != nil {
		// Pass the error up to the controller
		return models.User{}, err
	}

	return createdUser, nil
}
