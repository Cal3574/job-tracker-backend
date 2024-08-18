package joblog

import (
	"fmt"
	"job_tracker/internal/models/job_log"
	repositories "job_tracker/internal/repositories/job_log"
)

//Function to Create a new job log record
func CreateJobLog(title string, completed bool, note string, start_date string, end_date string, jobId int, categoryId int ) (job_log.JobLog, error) {
	jobLog := job_log.JobLog{
		Title: title,
		Completed: completed,
		Note: note,
		StartDate: start_date,
		EndDate: end_date,
		JobId: jobId,
		CategoryId: categoryId,

	}
	fmt.Println(jobLog.Title, "jobLog")
	return repositories.CreateJobLog(jobLog) 
}



//Function to find job log records by id
func FindJobLogById(jobId int) ([]job_log.JobLog, error) {
	return repositories.FindJobLogById(jobId)
}
