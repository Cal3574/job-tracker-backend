package industries

import (
	models "job_tracker/internal/models/industries"
	"job_tracker/pkg/utils"
)

// Func to get all industries
func GetAllIndustries() ([]models.Industry, error) {
	// Query the database
	rows, err := utils.DB.Query("SELECT id, name FROM industries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold industries
	var industries []models.Industry

	// Iterate over the rows
	for rows.Next() {
		var industry models.Industry
		// Scan the row into the industry struct
		if err := rows.Scan(&industry.ID, &industry.Name); err != nil {
			return nil, err
		}
		// Append the industry to the slice
		industries = append(industries, industry)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return industries, nil
}
