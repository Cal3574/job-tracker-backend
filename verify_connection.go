package main

import (
	"fmt"
	"log"

	"job_tracker/pkg/utils" // Adjust import path if necessary

	_ "github.com/lib/pq" // Import the pq driver
)

func main() {
	// Initialize the database connection
	utils.InitDB()

	// Check if the database is connected
	err := utils.DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}
