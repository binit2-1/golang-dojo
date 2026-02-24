package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		slog.Info("Incoming Request",
				  "Method", r.Method, 
				  "path", r.URL, 
				  "duration", time.Since(start).String(), 
				  "ip", r.RemoteAddr)
	})
}
