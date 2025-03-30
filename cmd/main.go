package main

import (
	"fmt"
	"time"

	"github.com/lmtani/learning-rate-limiter/internal/infra/storage"
	"github.com/lmtani/learning-rate-limiter/internal/infra/web"
	"github.com/lmtani/learning-rate-limiter/internal/infra/web/server"
	"github.com/lmtani/learning-rate-limiter/internal/usecase"
	"github.com/lmtani/learning-rate-limiter/pkg/limiter"
)

func main() {
	// Initialize Redis storage and check if is reachable
	redisStorage := storage.NewRedisStorage("localhost:6379", "", 0)
	if err := redisStorage.Client.Ping(redisStorage.Client.Context()).Err(); err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

	// Initialize rate limiter with Redis storage
	rateLimiter := limiter.NewRateLimiter(2, 10*time.Second, redisStorage)

	// Initialize web server
	webServer := server.NewWebServer(":8080", rateLimiter)

	// Add handlers to the web server
	helloUseCase := usecase.NewHelloUseCase()
	helloHandler := web.NewHelloHandler(helloUseCase)
	webServer.AddHandler("/hello", helloHandler.Handle)

	// Start the web server
	fmt.Println("Starting server on port 8080...")
	webServer.Start()
}
