package services

import (
	models "job_tracker/internal/models/job"
	repositories "job_tracker/internal/repositories/job"
)

//Function to find all jobs
func GetAllJobs(userId int) ([]models.Job, error) {
	return repositories.FindAllJobs(userId)
}

//Function to find a job by ID
func GetJobById(id int) (models.Job, error) {
	return repositories.GetJobById(id)
}


//Function to create a job
func CreateJob(title string, location string, company string, salary int, url string) (models.Job, error){
	job := models.Job{
		JobTitle: title,
		Location: location,
		Company: company,
		Salary: salary,
		URL: url,
	}
	return repositories.CreateJob(job)
}

//Function to delete a job
func DeleteJob(id int) error {
	return repositories.DeleteJob(id)
}

//Function to update a job
func UpdateJob(job models.Job) error {
	return repositories.UpdateJob(job)
}
