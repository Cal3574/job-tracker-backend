package interviews

import (
	"fmt"
	models "job_tracker/internal/models/interviews"
	repositories "job_tracker/internal/repositories/interviews"
	"job_tracker/pkg/email"
	"log"
	"time"
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
		t := time.Date(0000, 1, 1, 13, 0, 0, 0, time.UTC)
		readableTime := t.Format("15:04:05")

		// Send an email reminder to the user
		err := email.SendEmail(userEmail, name, "Interview Reminder", fmt.Sprintf("You have an upcoming interview tomorrow with %s at %s", interview.Company, readableTime))

		if err != nil {
			log.Fatalf("Error sending email reminder: %v", err)
		} else {
			// update the job log to indicate that the reminder has been sent
			repositories.UpdateInterviewReminderStatus(interview.JobLogID)
		}
	}
}
