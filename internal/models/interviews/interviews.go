package models

import (
	"database/sql"
	"time"
)

type UpcomingInterview struct {
	ID            int          `json:"id"`
	JobTitle      string       `json:"job_title"`
	Location      string       `json:"location"`
	Company       string       `json:"company"`
	Salary        float64      `json:"salary"`
	URL           string       `json:"url"`
	Priorities    string       `json:"priorities"`
	CreatedAt     time.Time    `json:"created_at"`
	UserID        int          `json:"user_id"`
	UserEmail     string       `json:"user_email"`
	UserFirstName string       `json:"user_first_name"`
	InterviewDate sql.NullTime `json:"interview_date"` // Nullable date
	InterviewTime sql.NullTime `json:"interview_time"` // Nullable time
	JobLogID      int          `json:"job_log_id"`
}

type InterviewReminder struct {
	JobLogID int `json:"job_log_id"`
}
