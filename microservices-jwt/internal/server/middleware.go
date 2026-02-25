package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/binit2-1/golang-dojo/microservices-jwt/pkg/jwtutil"
)

type ContextKey string

func (s *Server) RequireJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == authHeader {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := jwtutil.ValidateToken(token, s.JWTSecret)
		if err != nil {
			http.Error(w, "Invalid Token or Token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey("userEmail"), claims.Email)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
