package job_log

import (
	"fmt"
	models "job_tracker/internal/models/job_log"
	"job_tracker/pkg/utils"
	"log"
	"strconv"
)

// Function to create a new job log record
func CreateJobLog(jobLog models.JobLog) (models.JobLog, error) {
	log.Println(jobLog, "jobLog here!")
	fmt.Print(jobLog, "jobLog here!")
	var id int
	categoryId, _ := strconv.Atoi(jobLog.CategoryId)

	err := utils.DB.QueryRow("INSERT INTO job_log (title, complete, note, start_date, end_date, category_id, job_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", jobLog.Title, jobLog.Completed, jobLog.Note, jobLog.StartDate, jobLog.EndDate, categoryId, jobLog.JobId).Scan(&id)
	if err != nil {
		log.Println(err)
		return jobLog, err
	}

	fmt.Print(err, "err here!")
	jobLog.ID = id
	return jobLog, nil

}

// Function to find job log records by job ID
func FindJobLogById(jobId int) ([]models.JobLog, error) {
	rows, err := utils.DB.Query("SELECT id, title, complete, note, start_date, end_date, category_id, job_id FROM job_log WHERE job_id = $1", jobId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobLogs []models.JobLog
	for rows.Next() {
		var jobLog models.JobLog
		if err := rows.Scan(&jobLog.ID, &jobLog.Title, &jobLog.Completed, &jobLog.Note, &jobLog.StartDate, &jobLog.EndDate, &jobLog.CategoryId, &jobLog.JobId); err != nil {
			return nil, err
		}
		jobLogs = append(jobLogs, jobLog)
	}

	return jobLogs, nil
}

// Function to delete a job record by ID

func DeleteJobLogById(jobId int) error {
	_, err := utils.DB.Exec("DELETE FROM job_log WHERE id = $1", jobId)
	if err != nil {
		return err
	}
	return nil
}

// Function to update a job log record

func UpdateJobLog(jobLog models.JobLog) (models.JobLog, error) {
	// Debug: Print the jobLog ID
	fmt.Println("jobLog ID:", jobLog.ID)
	if jobLog.ID == 0 {
		fmt.Println("Error: jobLog ID is 0, no row will be updated")
		return jobLog, fmt.Errorf("invalid jobLog ID: %d", jobLog.ID)
	}

	// Convert category ID and debug print
	categoryId, err := strconv.Atoi(jobLog.CategoryId)
	if err != nil {
		fmt.Println("Error converting CategoryId:", jobLog.CategoryId, "Error:", err)
		return jobLog, err
	}
	fmt.Println("Converted CategoryId:", categoryId)

	fmt.Print(jobLog.ID, "jobLog here!")

	// Execute SQL query and debug print
	query := "UPDATE job_log SET title = $1, complete = $2, note = $3, start_date = $4, end_date = $5, category_id = $6, job_id = $7 WHERE id = $8"
	fmt.Println("Executing SQL Query:", query)
	fmt.Println("With Values:", jobLog.Title, jobLog.Completed, jobLog.Note, jobLog.StartDate, jobLog.EndDate, categoryId, jobLog.JobId, jobLog.ID)

	result, err := utils.DB.Exec(query, jobLog.Title, jobLog.Completed, jobLog.Note, jobLog.StartDate, jobLog.EndDate, categoryId, jobLog.JobId, jobLog.ID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return jobLog, err
	}

	// Debug: Check the result of the query
	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Rows affected:", rowsAffected)

	if rowsAffected == 0 {
		fmt.Println("Warning: No rows were updated. This might indicate an incorrect ID.")
	}

	fmt.Println("Updated JobLog:", jobLog)
	return jobLog, nil
}
