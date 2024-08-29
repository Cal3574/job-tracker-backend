package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const UserIDKey contextKey = "userId"

// JWTMiddleware is the middleware function for validating JWTs
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing or malformed Authorization header"))
			return
		}

		jwtToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			//! This is a secret key that should be stored securely TODO: Move this to a config file
			return []byte("123123123"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Token"))
			return
		}

		if !token.Valid {
			fmt.Println("Invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Token"))
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Token Claims"))
			return
		}

		// Extract userId from claims
		userId, ok := claims["userId"].(float64)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("User ID not found in token"))
			return
		}

		// Store the userId in the context
		ctx := context.WithValue(r.Context(), UserIDKey, int(userId))
		// Proceed to the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
