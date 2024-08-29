package joblog

import (
	"fmt"
	"job_tracker/internal/models/job_log"
	repositories "job_tracker/internal/repositories/job_log"
)

// Function to Create a new job log record
func CreateJobLog(title string, completed bool, note string, start_date *string, interview_date *string, jobId int, categoryId string, interview_time *string) (job_log.JobLog, error) {

	fmt.Println(title, completed, note, start_date, interview_date, jobId, categoryId)
	jobLog := job_log.JobLog{
		Title:         title,
		Completed:     completed,
		Note:          note,
		StartDate:     start_date,
		InterviewDate: interview_date,
		JobId:         jobId,
		CategoryId:    categoryId,
	}
	return repositories.CreateJobLog(jobLog)
}

// Function to find job log records by id
func FindJobLogById(jobId int) ([]job_log.JobLog, error) {
	return repositories.FindJobLogById(jobId)
}

// Function to delete a job log record by id
func DeleteJobLogById(jobId int) error {
	return repositories.DeleteJobLogById(jobId)
}

// Function to update a job log record
func UpdateJobLog(jobLog job_log.JobLog) (job_log.JobLog, error) {
	return repositories.UpdateJobLog(jobLog)
}
