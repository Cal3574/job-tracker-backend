package services

import (
	"database/sql"
	models "job_tracker/internal/models/users"
	"job_tracker/pkg/utils"
)

// Function to create a new user record or return an existing user
func CreateUser(user models.User) (models.User, bool, error) {
	// Check if the email already exists in the database
	existingUser, err := GetUserByEmail(user.Email)
	if err != nil && err != sql.ErrNoRows {
		// Return an error if something went wrong other than "no rows found"
		return user, false, err
	}

	if existingUser.Email != "" {
		// If an existing user is found with the email, return the existing user and false
		return existingUser, false, nil
	}

	// Insert the new user into the database
	var id int
	err = utils.DB.QueryRow("INSERT INTO users (email, name) VALUES ($1, $2) RETURNING id", user.Email, user.Name).Scan(&id)
	if err != nil {
		return user, false, err
	}

	// Assign the generated ID to the user model
	user.ID = id
	return user, true, nil // Return the new user and true indicating it's a new user
}

// Function to get a user by email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := utils.DB.QueryRow("SELECT id, email, name FROM users WHERE email=$1", email).Scan(&user.ID, &user.Email, &user.Name)

	if err != nil {
		// If no user is found, return an empty user and sql.ErrNoRows
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		// For other errors, return the error
		return user, err
	}

	return user, nil
}
