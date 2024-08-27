package industries

import (
	models "job_tracker/internal/models/industries"
	repositories "job_tracker/internal/repositories/industries"
)

// Func to get all industries
func GetAllIndustries() ([]models.Industry, error) {
	return repositories.GetAllIndustries()
}
