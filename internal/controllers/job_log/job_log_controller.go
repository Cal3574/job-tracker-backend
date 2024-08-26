package controllers

import (
	"encoding/json"
	"fmt"
	models "job_tracker/internal/models/job_log"
	services "job_tracker/internal/services/job_log"
	"net/http"
	"strconv"
)

// CreateJobLog handles POST requests to the /job_logs endpoint
// It creates a new job log record
func CreateJobLog(w http.ResponseWriter, r *http.Request) {
	var jobLog models.JobLog

	// Decode JSON request body into jobLog
	if err := json.NewDecoder(r.Body).Decode(&jobLog); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	// Prepare values for insertion
	var startDate, endDate *string
	if jobLog.StartDate != nil {
		startDate = jobLog.StartDate
	} else {
		startDate = nil
	}
	if jobLog.EndDate != nil {
		endDate = jobLog.EndDate
	} else {
		endDate = nil
	}

	// Call the service to create the job log
	newJobLog, err := services.CreateJobLog(jobLog.Title, jobLog.Completed, jobLog.Note, startDate, endDate, jobLog.JobId, jobLog.CategoryId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create job log: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the newly created job log
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newJobLog); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

//FindJobLogById handles GET requests to the /job_logs/{jobId} endpoint

func FindJobLogById(w http.ResponseWriter, r *http.Request) {
	jobId := r.URL.Query().Get("id")

	if jobId == "" {
		http.Error(w, "Missing jobId parameter", http.StatusBadRequest)
		return
	}

	//turn job id into an integer
	jobIdInt, err := strconv.Atoi(jobId)

	if err != nil {
		http.Error(w, "Invalid jobId parameter", http.StatusBadRequest)
		return
	}

	jobLog, err := services.FindJobLogById(jobIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(jobLog); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// DeleteJobLogById handles DELETE requests to the /job_logs/{jobId} endpoint

func DeleteJobLogById(w http.ResponseWriter, r *http.Request) {

	jobLogId := r.URL.Query().Get("id")

	//log for error handing

	fmt.Println(jobLogId, "jobLogId")

	if jobLogId == "" {
		http.Error(w, "Missing jobLogId parameter", http.StatusBadRequest)
		return
	}

	//turn job id into an integer
	jobIdInt, err := strconv.Atoi(jobLogId)

	if err != nil {
		http.Error(w, "Invalid jobLogId parameter", http.StatusBadRequest)
		return
	}

	err = services.DeleteJobLogById(jobIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdateJobLog handles PUT requests to the /job_logs endpoint
// It updates an existing job log record

func UpdateJobLog(w http.ResponseWriter, r *http.Request) {
	var jobLog models.JobLog

	// Decode JSON request body into jobLog
	if err := json.NewDecoder(r.Body).Decode(&jobLog); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Print(jobLog, "jobLog here!")

	// Log the job log details in a structured format
	fmt.Printf("Received Job Log:\n")
	fmt.Printf("Title: %s\n", jobLog.Title)
	fmt.Printf("Completed: %v\n", jobLog.Completed)
	fmt.Printf("Note: %s\n", jobLog.Note)
	fmt.Printf("Start Date: %s\n", jobLog.StartDate)
	fmt.Printf("End Date: %s\n", jobLog.EndDate)
	fmt.Printf("Job ID: %d\n", jobLog.JobId)
	fmt.Printf("Category ID: %d\n", jobLog.CategoryId)
	fmt.Print("\n")
	fmt.Printf("Job Log ID: %d\n", jobLog.ID)

	// Call the service to update the job log
	updatedJobLog, err := services.UpdateJobLog(jobLog)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update job log: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the updated job log
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedJobLog); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
