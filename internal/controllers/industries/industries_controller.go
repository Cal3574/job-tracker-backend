package industries

import (
	"encoding/json"
	"fmt"
	services "job_tracker/internal/services/industries"
	"net/http"
)

// GetAllIndustries handles GET requests to the /industries endpoint
// It returns a list of all industries
func GetAllIndustries(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetAllIndustries")
	// Example: Pass userId to the service layer if needed
	industries, err := services.GetAllIndustries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(industries)
}
