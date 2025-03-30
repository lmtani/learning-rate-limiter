package server

import (
	"net/http"

	"github.com/lmtani/learning-rate-limiter/pkg/limiter"
)

const TOO_MANY_REQUESTS = "you have reached the maximum number of requests or actions allowed within a certain time frame"

func RateLimitMiddleware(limiter *limiter.RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		key := r.Header.Get("API_KEY")

		if key == "" {
			key = ip
		}

		if !limiter.ShouldPass(key) {
			http.Error(w, TOO_MANY_REQUESTS, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
