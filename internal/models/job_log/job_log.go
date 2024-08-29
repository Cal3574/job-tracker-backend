package job_log

type JobLog struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Completed     bool    `json:"completed"`
	Note          string  `json:"note"`
	StartDate     *string `json:"start_date"`
	InterviewDate *string `json:"interview_date"`
	InterviewTime *string `json:"interview_time"`
	CategoryId    string  `json:"category_id"`
	JobId         int     `json:"job_id"`
	CreatedAt     string  `json:"created_at"`
}
