// pkg/utils/database.go

package utils

import (
	"database/sql"
	"fmt"
	"log"

	"job_tracker/internal/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func InitDB() {
    config.LoadConfig()
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    log.Println("Successfully connected to the database!")
}
