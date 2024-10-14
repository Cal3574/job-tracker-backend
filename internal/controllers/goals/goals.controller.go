package goals

import (
	"encoding/json"
	"fmt"
	models "job_tracker/internal/models/goals"
	services "job_tracker/internal/services/goals"
	"job_tracker/internal/shared"
	"log"
	"net/http"
	"strconv"
)

// CreateGoal handles POST requests to the /goals endpoint
// It creates a new goal
func CreateGoal(w http.ResponseWriter, r *http.Request) {

	user_id := r.URL.Query().Get("user_id")
	if user_id == "" {
		http.Error(w, "user_id not provided", http.StatusInternalServerError)
	}

	var goal models.Goal
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newJob, err := services.CreateGoal(user_id, goal.GoalName, goal.GoalDescription, goal.GoalAction, goal.GoalDeadline, goal.GoalTarget, goal.GoalType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newJob); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteGoal handles the DELETE requests to the /goal endpoint

func DeleteGoal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Missing goal ID", http.StatusBadRequest)
		return
	}

	idAsInt, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
	}

	err = services.DeleteGoal(idAsInt)

	if err != nil {
		log.Printf("Failed to delete goal: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetAllGoals handles GET requests to the /goals endpoint
func GetAllGoals(w http.ResponseWriter, r *http.Request) {

	user_id := r.URL.Query().Get("user_id")
	if user_id == "" {
		http.Error(w, "user_id not provided", http.StatusInternalServerError)
	}

	goals, err := services.GetAllGoals(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(goals); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SendGoalCompletion(w http.ResponseWriter, req *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Make sure to flush the headers right away
	w.(http.Flusher).Flush()

	// Create a channel to handle when the client closes the connection
	clientGone := req.Context().Done()

	// Listen to goal completion events and stream to the client
	for {
		select {
		case goalId := <-shared.GoalCompleteChannel:
			// Send goal completion event
			fmt.Fprintf(w, "event: goalComplete\n")
			fmt.Fprintf(w, "data: Goal %d achieved! Time to celebrate!\n\n", goalId)
			w.(http.Flusher).Flush() // Ensure data is sent immediately

		case <-clientGone:
			// Client disconnected, clean up
			fmt.Println("Client disconnected")
			return
		}
	}
}
