package analytics

import (
	"job_tracker/pkg/utils"
	"time"
)

// GetJobApplicationsCount retrieves job titles from the job table for a specific user within a given date range.
func GetJobApplicationsCount(userId string, startDate, endDate time.Time) ([]string, error) {
	query := `
        SELECT job_title
        FROM job
        WHERE user_id = $1
        AND created_at >= $2
        AND created_at < $3
    `
	rows, err := utils.DB.Query(query, userId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobTitles []string
	for rows.Next() {
		var jobTitle string
		if err := rows.Scan(&jobTitle); err != nil {
			return nil, err
		}
		jobTitles = append(jobTitles, jobTitle)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobTitles, nil
}
