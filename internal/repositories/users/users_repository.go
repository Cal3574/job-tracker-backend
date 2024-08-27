package services

import (
	"database/sql"
	"fmt"
	models "job_tracker/internal/models/users"
	"job_tracker/pkg/utils"
	"reflect"
	"strconv"
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
	err = utils.DB.QueryRow("INSERT INTO users (email, name, signup_complete) VALUES ($1, $2, $3) RETURNING id", user.Email, user.Name, false).Scan(&id)
	if err != nil {
		return user, false, err
	}

	// Assign the generated ID to the user model

	user.ID = strconv.Itoa(id)
	return user, true, nil // Return the new user and true indicating it's a new user
}

// CompleteSignUp function to update user details
func CompleteSignUp(user models.User) error {
	fmt.Println(reflect.TypeOf(user.ID), "type of user ID")
	fmt.Println(user.ID, "user ID")

	userId, err := strconv.Atoi(user.ID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}

	industryId, err := strconv.Atoi(user.DesiredIndustryId)
	if err != nil {
		return fmt.Errorf("invalid desired industry ID: %v", err)
	}

	isComplete := true

	query := `UPDATE users SET first_name=$1, last_name=$2, current_job_role=$3, experience_level=$4, desired_job_role=$5, desired_job_industry_id=$6, signup_complete=$7 WHERE id=$8`

	result, err := utils.DB.Exec(query,
		user.FirstName,
		user.LastName,
		user.CurrentJobRole,
		user.ExperienceLevel,
		user.DesiredJobRole,
		industryId,
		isComplete,
		userId,
	)

	if err != nil {
		return fmt.Errorf("failed to execute update query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated; possibly an invalid user ID")
	}

	return nil
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

// Function to check if a user has completed sign up process
func CheckUserSignUpStatus(userId int) (bool, error) {
	var signupComplete bool
	err := utils.DB.QueryRow("SELECT signup_complete FROM users WHERE id=$1", userId).Scan(&signupComplete)
	if err != nil {
		return false, err
	}

	return signupComplete, nil
}
