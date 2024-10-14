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

	// Fetch allowed origins from environment variable
	// Add support for multiple origins (local and production)
	allowedOrigins := []string{
		"http://localhost:3000",                                  // Local Development
		"https://job-tracker-frontend-production.up.railway.app", // Production Frontend
	}

	// Optionally, you can allow additional origins based on FRONTEND_URL env var
	if frontendURL := os.Getenv("FRONTEND_URL"); frontendURL != "" {
		allowedOrigins = append(allowedOrigins, frontendURL)
	}

	// Start the server with CORS middleware enabled
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)(router)))
}
