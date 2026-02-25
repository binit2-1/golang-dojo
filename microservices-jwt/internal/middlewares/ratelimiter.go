package middlewares

import (
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)


type RateLimiter struct{
	rdb *redis.Client
}

func NewRateLimiter(rdb *redis.Client)*RateLimiter{
	return &RateLimiter{rdb:rdb}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		key := "ratelimiter:login" + ip


		count, err := rl.rdb.Incr(r.Context(), key).Result()
		if err!=nil{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if count == 1{
			rl.rdb.Expire(r.Context(), key, 15*time.Minute)
		}

		if count > 15 {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
		}

		next.ServeHTTP(w, r)
	})
}