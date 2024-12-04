package middlewares

import (
	"net/http"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3) // 1 request/second, burst 3

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
