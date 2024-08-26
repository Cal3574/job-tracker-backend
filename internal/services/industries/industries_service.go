package industries

import (
	repositories "job_tracker/internal/repositories/industries"
	"job_tracker/internal/models/industries"
)

// Func to get all industries
func GetAllIndustries() (models., error) {
	return repositories.GetAllIndustries()
}
