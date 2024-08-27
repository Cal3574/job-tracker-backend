package models

//Composite data type that groups together related fields under a single name. It's similar to a object in JavaScript
//The 'json' tags tell GO how to encode/decode the struct when it serializes/deserializes it to/from JSON

type User struct {
	ID                string `json:"id"`
	CreatedAt         string `json:"created_at"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	CurrentJobRole    string `json:"job_role"`
	DesiredJobRole    string `json:"desired_job_role"`
	ExperienceLevel   string `json:"experience_level"`
	DesiredIndustryId string `json:"desired_job_industry_id"`
	SignupComplete    bool   `json:"signup_complete"`
}
