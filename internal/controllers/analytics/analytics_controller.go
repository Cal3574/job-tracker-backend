package industries

import (
	"encoding/json"
	"fmt"
	"job_tracker/internal/middleware"
	services "job_tracker/internal/services/analytics"
	"job_tracker/pkg/utils"
	"net/http"
	"time"
)

// GetApplicationCount handles GET requests to the /industries endpoint
// It returns a list of job titles for the last {days} days and the percentage change compared to the previous period
func GetApplicationCount(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	// Parse the 'days' parameter from the query string or set a default value
	days := 7
	queryDays := r.URL.Query().Get("days")
	if queryDays != "" {
		fmt.Sscanf(queryDays, "%d", &days)
		if days != 7 && days != 14 && days != 30 {
			http.Error(w, "Invalid days parameter. Must be 7, 14, or 30.", http.StatusBadRequest)
			return
		}
	}

	// Calculate the current period's start and end dates
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -days)

	// Calculate the previous period's start and end dates
	prevEndDate := startDate
	prevStartDate := prevEndDate.AddDate(0, 0, -days)

	// Retrieve job titles for the current and previous periods
	currentJobs, err := services.GetJobApplicationsCount(userId, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prevJobs, err := services.GetJobApplicationsCount(userId, prevStartDate, prevEndDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Calculate the percentage change
	currentCount := len(currentJobs)
	prevCount := len(prevJobs)
	percentageChange := utils.CalculatePercentageChange(currentCount, prevCount)

	// Prepare the response data
	responseData := map[string]interface{}{
		"jobTitles":        currentJobs,
		"percentageChange": percentageChange,
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}
