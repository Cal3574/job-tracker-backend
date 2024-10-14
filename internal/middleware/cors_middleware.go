package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// CORSHandler adds CORS headers to the response
func CORSHandler(next http.Handler) http.Handler {
	allowedOrigin := os.Getenv("FRONTEND_URL")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:3000" // Fallback if env var is not set
	}

	return handlers.CORS(
		handlers.AllowedOrigins([]string{allowedOrigin}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)(next)
}
