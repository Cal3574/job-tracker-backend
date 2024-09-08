package controllers

import (
	"encoding/json"
	"job_tracker/internal/middleware"
	models "job_tracker/internal/models/job"
	goal_services "job_tracker/internal/services/goals"
	services "job_tracker/internal/services/job"

	"log"
	"net/http"
	"strconv"
)

// GetJobs handles GET requests to the /jobs endpoint
// It returns a list of all jobs
func GetJobs(w http.ResponseWriter, r *http.Request) {
	// Retrieve userId from the context
	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	// Example: Pass userId to the service layer if needed
	jobs, err := services.GetAllJobs(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobs)
}

//GetJobById handles GET requests to the /jobs/{id} endpoint

//It returns the job with the specified ID

func GetJobById(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/jobs/"):]

	if idStr == "" {
		http.Error(w, "Missing job ID", http.StatusBadRequest)
		return
	}

	idAsInt, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	job, err := services.GetJobById(idAsInt)
	if err != nil {
		log.Printf("Failed to get job: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(job)
}

// CreateJob handles POST requests to the /jobs endpoint
// It creates a new job
func CreateJob(w http.ResponseWriter, r *http.Request) {

	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//handle where priroties is null
	newJob, err := services.CreateJob(job.JobTitle, job.Location, job.Company, job.Salary, job.URL, userId, job.Priorities)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Handle goal checking.
	// Check if the user has any goals for job applications.
	// If they do, check if the job application meets the goal criteria and progress goals if necessary.

	err = goal_services.HandleUserGoals(userId, "job_applications")
	if err != nil {
		log.Printf("Error handling user goals: %v", err)
	}

	if err := json.NewEncoder(w).Encode(newJob); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteJob handles DELETE requests to the /jobs/{id} endpoint
// It deletes the job with the specified ID
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing job ID", http.StatusBadRequest)
		return
	}

	idAsInt, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteJob(idAsInt)
	if err != nil {
		log.Printf("Failed to delete job: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

//UpdateJob handles PUT requests to the /jobs endpoint
//It updates the job with the specified ID

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.UpdateJob(job)
	if err != nil {
		log.Printf("Failed to update job: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Decoding: Converts JSON in the request body to a Go struct.
// Service Call: Creates a job in the backend and handles potential errors.
// Response Status: Sets the HTTP status code to 201 Created.
// Encoding: Converts the Go struct to JSON and writes it to the response body.
// Error Handling: Properly handles errors at each step, including JSON encoding.
