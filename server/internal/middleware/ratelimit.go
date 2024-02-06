package middleware

import (
	"net/http"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

func RateLimiterMiddleware(rps int, burst int) func(http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if limiter.Allow() {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusTooManyRequests)

				headers := w.Header()
				resetTime := time.Now().Add(time.Duration(limiter.Limit()) * time.Second)
				headers.Set("X-RateLimit-Limit", strconv.Itoa(rps))
				headers.Set("X-RateLimit-Remaining", strconv.Itoa(limiter.Burst()-1))
				headers.Set("X-RateLimit-Reset", resetTime.Format(time.RFC1123))

				return
			}
		})
	}
}
