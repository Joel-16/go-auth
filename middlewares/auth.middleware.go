package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"auth/config"
	"auth/models"

	"github.com/golang-jwt/jwt"
)

// Middleware to validate JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
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
			config.Respond(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			config.Respond(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}
		user_id, _ := strconv.ParseUint(userID, 10, 32)

		// Attach user_id to request context

		if err := config.DB.First(&user, "id=?", uint(user_id)).Error; err != nil {
			config.Respond(w, http.StatusBadRequest, "Invalid user")
			return
		}
		user.Password = ""
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
