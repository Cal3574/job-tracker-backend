package repositories

import (
	"fmt"
	models "job_tracker/internal/models/job"
	"job_tracker/pkg/utils"
)

// Function to find all jobs
// Returns a slice of Job structs
func FindAllJobs(userId int) ([]models.Job, error) {
    fmt.Println("User ID: ", userId)
    // Execute a query to retrieve all rows from the 'job' table
    rows, err := utils.DB.Query("SELECT id, job_title, company, location, salary, url FROM job WHERE user_id = $1", userId)
    if err != nil {
        // If there is an error executing the query, return nil and the error
        return nil, err
    }
    // Ensure that the rows are closed when the function exits
	// This closes the result set and releases any associated database resources.
	// Defer ensures it is called as the last statement in the function
    defer rows.Close()

    // Create a slice to hold the job records
    var jobs []models.Job
    // Iterate over the result set
	
    for rows.Next() {
        // Create a variable to hold the data for each row
        var job models.Job
        // Scan the columns of the current row into the 'job' variable
				
		//Pointers to Variables: &job.ID, &job.JobTitle, etc., are pointers to the fields of the job struct. Scan will use these pointers to write the data from the row into these fields.
		// Column Matching: Scan expects the order of pointers to match the order of columns returned by the SQL query. In this case, it expects id, job_title, company, location, salary, url in that order.
        if err := rows.Scan(&job.ID, &job.JobTitle, &job.Company, &job.Location, &job.Salary, &job.URL); err != nil {
            // If there is an error scanning the row, return nil and the error
            return nil, err
        }
        // Append the job to the slice
        jobs = append(jobs, job)
    }
    // Return the slice of jobs and a nil error
    return jobs, nil
}

func GetJobById(id int) (models.Job, error) {
    fmt.Println("ID: ", id)
    var job models.Job
    err := utils.DB.QueryRow("SELECT id, job_title, company, location, salary, url FROM job WHERE id = $1", id).Scan(&job.ID, &job.JobTitle, &job.Company, &job.Location, &job.Salary, &job.URL)
    return job, err
}



//Function to create a job
//Takes a Job struct as an argument
//mutex.Lock() locks the mutex so that only one goroutine can access the jobs slice at a time
//The defer keyword is used to defer the execution of the mutex.Unlock() statement until the function returns
//The mutex is unlocked
func CreateJob(job models.Job) (models.Job, error) {
    var id int
    err := utils.DB.QueryRow(
        `INSERT INTO job(job_title, company, location, salary, url ) VALUES($1, $2, $3, $4, $5) RETURNING id`,
        job.JobTitle, job.Company, job.Location, job.Salary, job.URL,
    ).Scan(&id)
    if err != nil {
        return job, err
    }
    job.ID = id
    return job, nil
}

//Function to delete a job by ID
//Takes an integer ID as an argument
//Executes a SQL query to delete the job with the specified ID
func DeleteJob(id int) error {
	_, err := utils.DB.Exec("DELETE FROM job WHERE id = $1", id)
	return err
}


//Function to update a job
//Takes a Job struct as an argument
//Executes a SQL query to update the job with the specified ID
func UpdateJob(job models.Job) error {
		_, err := utils.DB.Exec(
			"UPDATE job SET job_title = $1, company = $2, location = $3, salary = $4, url = $5 WHERE id = $6",
			job.JobTitle, job.Company, job.Location, job.Salary, job.URL, job.ID,
		)
		return err
	}