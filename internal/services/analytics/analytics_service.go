package analytics

import (
	repositories "job_tracker/internal/repositories/analytics"
	"time"
)

// GetJobApplicationsCount retrieves job titles for a user within the specified date range.
func GetJobApplicationsCount(userId int, startDate, endDate time.Time) ([]string, error) {
	return repositories.GetJobApplicationsCount(userId, startDate, endDate)
}
