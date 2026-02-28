package server

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var clients = make(map[string]*rate.Limiter)
var mu sync.RWMutex

func RateLimiterMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		mu.RLock()
		limiter, exists := clients[ip]
		mu.RUnlock()
		
		if !exists{
			mu.Lock()
			limiter, exists = clients[ip]

			if !exists{
				limiter = rate.NewLimiter(2, 4)
				clients[ip] = limiter
			}

			mu.Unlock()
		} 

		if !limiter.Allow(){
			http.Error(w, "429 Too Many Requests - Slow down!", http.StatusTooManyRequests)
			return 
		}

		next.ServeHTTP(w, r)
	})
}

func MetricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Seconds()
		
		http_requests_total.Inc()
		http_request_duration_seconds.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}