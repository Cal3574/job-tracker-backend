package repositories

import (
	"fmt"
	models "job_tracker/internal/models/job"
	"job_tracker/pkg/utils"
)

// Function to find all jobs
// Returns a slice of Job structs
func FindAllJobs(userId int) ([]models.Job, error) {
	rows, err := utils.DB.Query(`
        SELECT 
            j.id, 
            j.job_title, 
            j.company, 
            j.location, 
            j.salary, 
            j.url, 
            j.priorities, 
            jsc.category_name,
			j.created_at
        FROM 
            job j
        LEFT JOIN 
            job_log jl ON j.id = jl.job_id AND jl.created_at = (
                SELECT MAX(created_at) 
                FROM job_log 
                WHERE job_id = j.id
            )
        LEFT JOIN 
            job_status_categories jsc ON jl.category_id = jsc.id
        WHERE 
            j.user_id = $1
        ORDER BY 
            j.id`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job

	for rows.Next() {
		var job models.Job

		// Use sql.NullString for CategoryName to handle NULL values
		if err := rows.Scan(&job.ID, &job.JobTitle, &job.Company, &job.Location, &job.Salary, &job.URL, &job.Priorities, &job.CategoryName, &job.CreatedAt); err != nil {
			return nil, err
		}

		// Append the job to the slice
		jobs = append(jobs, job)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}
func GetJobById(id int) (models.Job, error) {
	var job models.Job
	err := utils.DB.QueryRow("SELECT id, job_title, company, location, salary, url, priorities FROM job WHERE id = $1", id).Scan(&job.ID, &job.JobTitle, &job.Company, &job.Location, &job.Salary, &job.URL, &job.Priorities)
	return job, err
}

// Function to create a job
// Takes a Job struct as an argument
// mutex.Lock() locks the mutex so that only one goroutine can access the jobs slice at a time
// The defer keyword is used to defer the execution of the mutex.Unlock() statement until the function returns
// The mutex is unlocked
func CreateJob(job models.Job, userId int) (models.Job, error) {
	var id int
	err := utils.DB.QueryRow(
		`INSERT INTO job(job_title, company, location, salary, url, user_id, priorities ) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		job.JobTitle, job.Company, job.Location, job.Salary, job.URL, userId, job.Priorities,
	).Scan(&id)

	if err != nil {
		return job, err
	}
	job.ID = id
	return job, nil
}

// Function to delete a job by ID
// Takes an integer ID as an argument
// Executes a SQL query to delete the job with the specified ID
func DeleteJob(id int) error {
	_, err := utils.DB.Exec("DELETE FROM job WHERE id = $1", id)
	return err
}

// Function to update a job
// Takes a Job struct as an argument
// Executes a SQL query to update the job with the specified ID
func UpdateJob(job models.Job) error {
	result, err := utils.DB.Exec(
		"UPDATE job SET job_title = $1, company = $2, location = $3, salary = $4, url = $5, priorities = $7 WHERE id = $6",
		job.JobTitle, job.Company, job.Location, job.Salary, job.URL, job.ID, job.Priorities,
	)
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
