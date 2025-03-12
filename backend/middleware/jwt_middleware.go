// backend/middleware/jwt_middleware.go
package middleware

import (
	"log"
	"net/http"
	"strings"

	"backend/utils"
)

// JWTAuthMiddleware protects routes requiring authentication
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware triggered: Checking Authorization Header...")

		authHeader := r.Header.Get("Authorization")
		log.Println("Authorization Header:", authHeader) // Debugging log

		if authHeader == "" {
			log.Println("ERROR: Authorization header missing")
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Expect "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // No "Bearer " prefix
			log.Println("ERROR: Invalid token format")
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		_, err := utils.ValidateJWT(tokenString)
		if err != nil {
			log.Println("ERROR: Invalid or expired token:", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		log.Println("Token is valid, proceeding to next handler")
		next.ServeHTTP(w, r)
	})
}
