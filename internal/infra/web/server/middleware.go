package server

import (
	"net/http"

	"github.com/lmtani/learning-rate-limiter/pkg/limiter"
)

const TOO_MANY_REQUESTS = "you have reached the maximum number of requests or actions allowed within a certain time frame"

func RateLimitMiddleware(limiter *limiter.RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract API key from header or use IP as fallback
		var key, limiterType string
		if apiKey := r.Header.Get("X-API-Key"); apiKey != "" {
			key = apiKey
			limiterType = "api_key"
		} else {
			key = r.RemoteAddr
			limiterType = "ip"
		}

		if !limiter.ShouldPass(key, limiterType) {
			http.Error(w, TOO_MANY_REQUESTS, http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
