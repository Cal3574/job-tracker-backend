package services

import (
	"database/sql"
	"fmt"
	models "job_tracker/internal/models/users"
	"job_tracker/pkg/utils"
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
	err = utils.DB.QueryRow("INSERT INTO users (email, name, signup_complete, user_id) VALUES ($1, $2, $3, $4) RETURNING id", user.Email, user.Name, false, user.UserId).Scan(&id)
	if err != nil {
		return user, false, err
	}

	// Assign the generated ID to the user model

	user.ID = strconv.Itoa(id)
	return user, true, nil // Return the new user and true indicating it's a new user
}

// CompleteSignUp function to update user details
func CompleteSignUp(user models.User) error {

	industryId, err := strconv.Atoi(user.DesiredIndustryId)
	if err != nil {
		return fmt.Errorf("invalid desired industry ID: %v", err)
	}

	isComplete := true

	query := `UPDATE users SET first_name=$1, last_name=$2, current_job_role=$3, experience_level=$4, desired_job_role=$5, desired_job_industry_id=$6, signup_complete=$7 WHERE user_id=$8`

	result, err := utils.DB.Exec(query,
		user.FirstName,
		user.LastName,
		user.CurrentJobRole,
		user.ExperienceLevel,
		user.DesiredJobRole,
		industryId,
		isComplete,
		user.UserId,
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
func CheckUserSignUpStatus(userId string) (bool, error) {
	var signupComplete bool
	err := utils.DB.QueryRow("SELECT signup_complete FROM users WHERE user_id=$1", userId).Scan(&signupComplete)
	if err != nil {
		return false, err
	}

	return signupComplete, nil
}

// Function to get the users personal information
func GetUserInformation(userId string) (models.UserInfo, error) {
	var user_info models.UserInfo
	fmt.Println(userId, "userId here")
	err := utils.DB.QueryRow("SELECT users.id, users.first_name, users.last_name, users.email, users.current_job_role, users.desired_job_role, users.experience_level, industries.name FROM users JOIN industries ON users.desired_job_industry_id = industries.id WHERE users.user_id=$1", userId).Scan(
		&user_info.ID,
		&user_info.FirstName,
		&user_info.LastName,
		&user_info.Email,
		&user_info.CurrentJobRole,
		&user_info.DesiredJobRole,
		&user_info.ExperienceLevel,
		&user_info.DesiredJobIndustry,
	)

	fmt.Println(user_info, "user here")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Printing line 122, no rows found")
			return models.UserInfo{}, nil
		}
		// For other errors, return the error
		fmt.Println(err, "err here in getuserinformaiton")
		fmt.Println("Printing line 127, error")

		return user_info, err
	}
	fmt.Println("Printing line 130, user found")

	return user_info, nil
}

// Func to update user personal details
func UpdateUserPersonalDetails(user_personal_info models.UserPersonalInfo) error {
	fmt.Println(user_personal_info, "user personal info here")
	result, err := utils.DB.Exec("UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE user_id = $4", user_personal_info.FirstName, user_personal_info.LastName, user_personal_info.Email, user_personal_info.ID)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println("Error fetching rows affected:", err)

	} else {
		fmt.Println("Rows affected:", rowsAffected)

	}
	return nil
}

// Func to update user career details
func UpdateUserCareerDetails(user_personal_info models.UserCareerInfo) error {
	fmt.Println(user_personal_info, "user career info here")
	result, err := utils.DB.Exec("UPDATE users SET current_job_role = $1, experience_level = $2, desired_job_role = $3, desired_job_industry_id = $4 WHERE user_id = $5", user_personal_info.CurrentJobRole, user_personal_info.ExperienceLevel, user_personal_info.DesiredJobRole, user_personal_info.DesiredIndustryId, user_personal_info.ID)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println("Error fetching rows affected:", err)

	} else {
		fmt.Println("Rows affected:", rowsAffected)

	}
	return nil
}
