package middleware

import (
	"net/http"
	"sync"
)

type RateLimiter struct{
	visitors map[string]int
	mu 		 sync.Mutex
}

func NewRateLimiter() *RateLimiter{
	
	return &RateLimiter{
		visitors: make(map[string]int), //can't use maps without initializing
		//sync mutex is zero valued you can use almost immediately as it is intialized as zero
	}
}

func (rl *RateLimiter) Limit(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ip := r.RemoteAddr //In production behind a load balancer, you'd check the X-Forwarded-For header, but RemoteAddr works for local testing
		//lock as soon as new request comes
		rl.mu.Lock()
		
		//checks if the same ip has made more than 5 requests
		if rl.visitors[ip] >= 5{
			//unlock before returning to allow other requests to proceed
			rl.mu.Unlock()
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			return
		}
		
		rl.visitors[ip]++
		rl.mu.Unlock() //unlock after updating the count
		next.ServeHTTP(w, r)
	})
}