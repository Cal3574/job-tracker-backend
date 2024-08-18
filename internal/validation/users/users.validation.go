package users

import "regexp"

// isValidEmail checks if the provided email has a valid format
func IsValidEmail(email string) bool {
	// A simple regex for basic email validation
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}