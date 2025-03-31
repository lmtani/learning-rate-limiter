package server

import (
	"net/http"

	"github.com/lmtani/learning-rate-limiter/pkg/limiter"
)

type WebServer struct {
	Router        *http.ServeMux
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
	RateLimiter   *limiter.RateLimiter
}

func NewWebServer(serverPort string, rateLimiter *limiter.RateLimiter) *WebServer {
	return &WebServer{
		Router:        http.NewServeMux(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
		RateLimiter:   rateLimiter,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}

	// Use Request Logging Middleware
	loggedRouter := LoggingMiddleware(s.Router)

	// Use limiter middleware
	limitedRouter := RateLimitMiddleware(s.RateLimiter, loggedRouter)

	err := http.ListenAndServe(s.WebServerPort, limitedRouter)
	if err != nil {
		panic(err)
	}
}
