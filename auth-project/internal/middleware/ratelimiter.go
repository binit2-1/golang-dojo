package middleware

import (
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)




type RateLimiter struct{
	rdb *redis.Client
}

func NewRateLimiter(rdb *redis.Client)*RateLimiter{
	return &RateLimiter{rdb: rdb}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		ip := r.RemoteAddr
		key :="ratelimit:login" + ip

		// 1. Increment the count for this IP
		// INCR creates the key with value 1 if it doesn't exist
		count, err := rl.rdb.Incr(r.Context(), key).Result()
		if err!=nil{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// 2. If it's the first request, set the window to 15 minutes
		if count == 1{
			rl.rdb.Expire(r.Context(), key, 15*time.Minute)
		}

		// 3. Check if they exceeded the limit (e.g., 5 attempts)
		if count >5{
			http.Error(w, "Too Many Requests. Try again in 15 minutes.", http.StatusTooManyRequests)
			return 
		}

		next.ServeHTTP(w, r)
	})
}