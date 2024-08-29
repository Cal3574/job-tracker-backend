package interviews

import (
	"fmt"
	models "job_tracker/internal/models/interviews"
	repositories "job_tracker/internal/repositories/interviews"
	"job_tracker/pkg/email"
)

// Function to find all jobs
func GetUpcomingInterviews() ([]models.UpcomingInterview, error) {
	return repositories.GetUpcomingInterviews()
}

// Function to send a reminder to the user via EMAIL
func SendReminderEmail(interview_reminders []models.UpcomingInterview) {
	// Send email reminders to the user

	// Iterate over the list of upcoming interviews
	for _, interview := range interview_reminders {

		//destructure the interview object
		userEmail := interview.UserEmail
		name := interview.UserFirstName

		// Send an email reminder to the user
		email.SendEmail(userEmail, name, "Interview Reminder", fmt.Sprintf("You have an upcoming interview tomorrow with %s", interview.Company))

	}
}
