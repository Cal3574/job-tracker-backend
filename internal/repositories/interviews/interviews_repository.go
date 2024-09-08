package interviews

import (
	models "job_tracker/internal/models/interviews"
	"job_tracker/pkg/utils"
	"log"
)

// Func to search for all jobs that have upcoming interviews
func GetUpcomingInterviews() ([]models.UpcomingInterview, error) {
	query := `
        SELECT 
            j.id, 
            j.job_title, 
            j.location, 
            j.company, 
            j.salary, 
            j.url, 
            j.priorities, 
            j.created_at,
            j.user_id,
            u.email AS user_email,
			u.first_name AS user_first_name,
            jl.interview_date,
            jl.interview_time,
			jl.id AS job_log_id
        FROM 
            job_log AS jl
        JOIN 
            job AS j 
        ON 
            jl.job_id = j.id
        JOIN 
            users AS u
        ON
            j.user_id = u.id
        WHERE 
            jl.category_id = 5
			AND
			(jl.interview_reminder = false OR jl.interview_reminder IS NULL)        
			AND 
            (jl.interview_date + jl.interview_time::time) 
            BETWEEN NOW() AND (NOW() + INTERVAL '24 hours');
    `

	rows, err := utils.DB.Query(query)

	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var jobs []models.UpcomingInterview

	// Iterate over the result set
	for rows.Next() {
		var jobs_with_interviews models.UpcomingInterview
		err := rows.Scan(
			&jobs_with_interviews.ID,
			&jobs_with_interviews.JobTitle,
			&jobs_with_interviews.Location,
			&jobs_with_interviews.Company,
			&jobs_with_interviews.Salary,
			&jobs_with_interviews.URL,
			&jobs_with_interviews.Priorities,
			&jobs_with_interviews.CreatedAt,
			&jobs_with_interviews.UserID,
			&jobs_with_interviews.UserEmail,
			&jobs_with_interviews.UserFirstName,
			&jobs_with_interviews.InterviewDate,
			&jobs_with_interviews.InterviewTime,
			&jobs_with_interviews.JobLogID,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		jobs = append(jobs, jobs_with_interviews)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Printf("Error during row iteration: %v", err)
		return nil, err
	}

	return jobs, nil
}

// Func to update the interview reminder status in job_log table
func UpdateInterviewReminderStatus(jobLogID int) error {
	query := `
		UPDATE 
			job_log
		SET 
			interview_reminder = true
		WHERE 
			id = $1;
	`
	_, err := utils.DB.Exec(query, jobLogID)

	if err != nil {
		log.Printf("Error updating interview reminder status: %v", err)
		return err
	}

	return nil
}
