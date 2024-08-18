package job_log

import (
	models "job_tracker/internal/models/job_log"
	"job_tracker/pkg/utils"
	"log"
)

//Function to create a new job log record
func CreateJobLog(jobLog models.JobLog) (models.JobLog, error) {
log.Println(jobLog)
	var id int
	err := utils.DB.QueryRow("INSERT INTO job_log (title, complete, note, start_date, end_date, category_id, job_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", jobLog.Title, jobLog.Completed, jobLog.Note, jobLog.StartDate, jobLog.EndDate, jobLog.CategoryId, jobLog.JobId).Scan(&id)
	if err != nil {
		log.Println(err)
		return jobLog, err
	}

	jobLog.ID = id
	return jobLog, nil
	
}

//Function to find job log records by job ID
func FindJobLogById(jobId int) ([]models.JobLog,error) {
	rows, err := utils.DB.Query("SELECT id, title, complete, note, start_date, end_date, category_id, job_id FROM job_log WHERE jobId = $1", jobId)
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