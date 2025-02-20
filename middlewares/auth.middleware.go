package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"auth/config"

	"github.com/golang-jwt/jwt"
)

// Middleware to validate JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Parse token
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}
		user_id, _ := strconv.ParseUint(userID, 10, 32)

		// Attach user_id to request context
		ctx := context.WithValue(r.Context(), "user_id", uint(user_id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
