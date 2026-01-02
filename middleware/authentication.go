package middleware
import (
	"fmt"
	"net/http"
	"os"
	"github.com/golang-jwt/jwt/v5"
)
var jwtSecret []byte
// Middleware to verify JWT token
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		var tokenString string
		fmt.Sscanf(authHeader, "Bearer %s", &tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token valid → call next handler
		next.ServeHTTP(w, r)
	}
}