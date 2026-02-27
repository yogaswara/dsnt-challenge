package http

import (
	"net/http"
	"strings"

	"dsnt-challenge/internal/core/ports"
	"dsnt-challenge/pkg/logger"
)

// AuthMiddleware intercepts requests and asserts valid Bearer tokens
func AuthMiddleware(authService ports.AuthService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logger.Error("missing authorization header", nil)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			logger.Error("invalid authorization header format", nil)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := parts[1]
		if err := authService.ValidateToken(r.Context(), token); err != nil {
			logger.Error("invalid token", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed context
		next.ServeHTTP(w, r)
	}
}
