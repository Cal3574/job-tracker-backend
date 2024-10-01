// routes/routes.go
package routes

import (
	analytics_controllers "job_tracker/internal/controllers/analytics"
	goal_controllers "job_tracker/internal/controllers/goals"
	industries_controllers "job_tracker/internal/controllers/industries"
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
	authRouter.HandleFunc("/analytics/application_count", analytics_controllers.GetApplicationCount).Methods("GET")
	authRouter.HandleFunc("/goals", goal_controllers.CreateGoal).Methods("POST")
	authRouter.HandleFunc("/goals", goal_controllers.GetAllGoals).Methods("GET")
	authRouter.HandleFunc("/goals", goal_controllers.DeleteGoal).Methods("DELETE")
	authRouter.HandleFunc("/users/user_data", user_controllers.GetUserInformation).Methods("GET")
	authRouter.HandleFunc("/users/user_data/personal", user_controllers.UpdateUserPersonalDetails).Methods("PUT")
	authRouter.HandleFunc("/users/user_data/career", user_controllers.UpdateUserCareerDetails).Methods("PUT")

	// Define routes for users that do not need authentication
	router.HandleFunc("/users", user_controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", user_controllers.CompleteSignUp).Methods("PUT")
	router.HandleFunc("/industries", industries_controllers.GetAllIndustries).Methods("GET")
	router.HandleFunc("/users/check_signup_status", user_controllers.CheckUserSignUpStatus).Methods("GET")

	//SHOULD THIS BE AUTH ROUTE OR NOT?
	router.HandleFunc("/goals/completion", goal_controllers.SendGoalCompletion).Methods("GET")
	return router
}
