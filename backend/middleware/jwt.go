package middleware

import (
	"context"
	"net/http"
	"strings"

	"go_health_sentiment/auth"
	"go_health_sentiment/utils"
)

type key int

const UserKey key = 0

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}