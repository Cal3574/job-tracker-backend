package cron

import (
	"fmt"
	"job_tracker/internal/services/interviews"
	"log"

	"github.com/robfig/cron/v3"
)

func StartCronJobs() {
	c := cron.New(cron.WithSeconds())

	// Schedule the job at a specific time (e.g., 9:00 AM every day)
	_, err := c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("CRON job executed at 9:00 AM")
		// GetUpcomingInterviews()

		// Call the function to get upcoming interviews
		upcomingInterviews, err := interviews.GetUpcomingInterviews()

		// test call email API
		// email.SendEmail(upcomingInterviews[0].UserEmail, "Callum", "Interview Reminder", fmt.Sprintf("You have an upcoming interview tomorrow at %s", upcomingInterviews[0].InterviewTime))

		// Send email reminders to the user
		interviews.SendReminderEmail(upcomingInterviews)

		if err != nil {
			log.Fatalf("Error getting upcoming interviews: %v", err)
		}
		fmt.Println("Upcoming interviews: ", upcomingInterviews)
	})

	if err != nil {
		log.Fatalf("Error scheduling CRON job: %v", err)
	}

	// Start the CRON scheduler
	c.Start()
}
