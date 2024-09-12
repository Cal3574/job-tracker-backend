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

// CompleteSignUp completes the user sign up process.
// Adding additional user information to the user record.
func CompleteSignUp(user models.User) error {
	err := repositories.CompleteSignUp(user)
	if err != nil {
		return err
	}
	return nil
}

// CheckUserSignUpStatus checks if the user has completed the sign up process.
func CheckUserSignUpStatus(userId int) (bool, error) {
	signUpStatus, err := repositories.CheckUserSignUpStatus(userId)
	if err != nil {
		return false, err
	}
	return signUpStatus, nil
}

// GetUserInformation to grab all the user information for profile page
func GetUserInformation(userId int) (models.UserInfo, error) {
	userInfo, err := repositories.GetUserInformation(userId)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

// Update user personal information firstname, lastname, email
func UpdateUserPersonalDetails(user_personal_info models.UserPersonalInfo) error {
	return repositories.UpdateUserPersonalDetails(user_personal_info)
}

// Update user career information job_role, experience_level, desired_job_role, desired_job_industry_id
func UpdateUserCareerDetails(user_personal_info models.UserCareerInfo) error {
	return repositories.UpdateUserCareerDetails(user_personal_info)
}
