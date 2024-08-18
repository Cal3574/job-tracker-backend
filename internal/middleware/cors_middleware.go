// middleware/cors.go
package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSHandler adds CORS headers to the response
func CORSHandler(next http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Replace with your frontend URL
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)(next)
}
