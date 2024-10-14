package main

import (
	"job_tracker/internal/routes"
	"job_tracker/pkg/utils"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
)

func main() {
	// Initialize the database connection
	log.Println("Initializing database connection...")
	utils.InitDB()

	// Set up routes
	log.Println("Setting up routes...")
	router := routes.SetupRoutes()
	log.Println("Routes set up successfully.")

	// Fetch allowed origins from environment variable
	allowedOrigins := []string{
		"http://localhost:3000",                                  // Local Development
		"https://job-tracker-frontend-production.up.railway.app", // Production Frontend
	}

	// Check if FRONTEND_URL environment variable is set
	if frontendURL := os.Getenv("FRONTEND_URL"); frontendURL != "" {
		log.Printf("Adding allowed origin from environment: %s\n", frontendURL)
		allowedOrigins = append(allowedOrigins, frontendURL)
	} else {
		log.Println("FRONTEND_URL not set, using defaults.")
	}

	// Log the allowed origins for debugging purposes
	log.Printf("Allowed Origins: %v\n", allowedOrigins)

	// Dynamic port assignment
	portStr := os.Getenv("PORT")
	var port int
	var err error

	if portStr != "" {
		port, err = strconv.Atoi(portStr) // Convert string to int
		if err != nil {
			log.Fatalf("Invalid PORT value: %s\n", portStr)
		}
	} else {
		port = 3000 // Default port
	}

	// Start the server with CORS middleware enabled
	log.Printf("Starting server on port %d...\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowCredentials(),
	)(router)))
}
