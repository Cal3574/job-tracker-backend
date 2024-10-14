// cmd/main.go

package main

import (
	"job_tracker/internal/routes"
	"job_tracker/pkg/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	// Initialize the database connection
	utils.InitDB()

	// Start the CRON jobs
	// cron.StartCronJobs()

	// Set up routes
	router := routes.SetupRoutes()

	allowedOrigin := os.Getenv("FRONTEND_URL")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:3000" // Fallback if env var is not set
	}

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{allowedOrigin}), // Replace with your frontend URL
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)(router)))

}
