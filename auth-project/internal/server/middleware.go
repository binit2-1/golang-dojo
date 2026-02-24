package server

import (
	"context"
	"net/http"
)

type ContextKey string

func(s *Server) RequireAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err!=nil{
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		email, err := s.rdb.Get(r.Context(), cookie.Value).Result()
		if err!=nil{
			http.Error(w, "Session Expired", http.StatusUnauthorized)
			return 
		}

		ctx := context.WithValue(r.Context(), ContextKey("userEmail"), email)
		reqWithCtx := r.WithContext(ctx)

		next.ServeHTTP(w, reqWithCtx)
	})
}