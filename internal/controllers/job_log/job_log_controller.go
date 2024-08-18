package controllers

import (
	"encoding/json"
	models "job_tracker/internal/models/job_log"
	services "job_tracker/internal/services/job_log"
	"net/http"
)

//CreateJobLog handles POST requests to the /job_logs endpoint
//It creates a new job log record
func CreateJobLog(w http.ResponseWriter, r *http.Request) {
	var job_log models.JobLog
	err := json.NewDecoder(r.Body).Decode(&job_log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newJobLog, err := services.CreateJobLog(job_log.Title, job_log.Completed, job_log.Note, job_log.StartDate, job_log.EndDate, job_log.JobId, job_log.CategoryId)

	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newJobLog); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

//FindJobLogById handles GET requests to the /job_logs/{jobId} endpoint