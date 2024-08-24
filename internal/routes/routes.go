// routes/routes.go
package routes

import (
	job_controllers "job_tracker/internal/controllers/job"
	job_log_controllers "job_tracker/internal/controllers/job_log"
	user_controllers "job_tracker/internal/controllers/users"
	"job_tracker/internal/middleware"

	"github.com/gorilla/mux"
)

// SetupRoutes configures the HTTP routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Apply CORS middleware to the router

	// Create a subrouter for routes that need authentication
	authRouter := router.PathPrefix("/").Subrouter()

	// Apply JWT middleware to the subrouter
	authRouter.Use(middleware.JWTMiddleware)

	// Define routes for jobs under the authenticated subrouter
	authRouter.HandleFunc("/jobs", job_controllers.GetJobs).Methods("GET")
	authRouter.HandleFunc("/jobs", job_controllers.GetJobById).Methods("GET")
	authRouter.HandleFunc("/jobs", job_controllers.CreateJob).Methods("POST")
	authRouter.HandleFunc("/jobs", job_controllers.DeleteJob).Methods("DELETE")
	authRouter.HandleFunc("/jobs", job_controllers.UpdateJob).Methods("PUT")
	authRouter.HandleFunc("/job_logs", job_log_controllers.CreateJobLog).Methods("POST")
	authRouter.HandleFunc("/job_logs", job_log_controllers.FindJobLogById).Methods("GET")
	authRouter.HandleFunc("/job_logs", job_log_controllers.DeleteJobLogById).Methods("DELETE")
	authRouter.HandleFunc("/job_logs", job_log_controllers.UpdateJobLog).Methods("PUT")

	// Define routes for users that do not need authentication
	router.HandleFunc("/users", user_controllers.CreateUser).Methods("POST")

	return router
}
