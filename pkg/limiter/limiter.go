package limiter

import (
	"fmt"
	"time"

	"github.com/lmtani/learning-rate-limiter/internal/entity"
)

type RateLimiterI interface {
	ShouldPass(key string) bool
}

type RateLimiterStore interface {
	Increment(key string, expire time.Duration) (int, error)
	Reset(key string) error
}

type RateLimiter struct {
	Limit    int
	Expire   time.Duration
	Store    RateLimiterStore
	TokenMap entity.TokenMap
}

func NewRateLimiter(limit int, expire time.Duration, store RateLimiterStore, tokenMap entity.TokenMap) *RateLimiter {
	return &RateLimiter{
		Limit:    limit,
		Expire:   expire,
		Store:    store,
		TokenMap: tokenMap,
	}
}

func (rl *RateLimiter) ShouldPass(key string, limitType string) bool {
	if limitType != "api_key" && limitType != "ip" {
		fmt.Println("Invalid limit type. Use 'api_key' or 'ip'.")
		return false
	}

	limit := rl.Limit // Na falta de api_key, limite padrão é mesmo utilizado para IP
	if limitType == "api_key" {
		if tokenLimit, exists := rl.TokenMap[key]; exists {
			limit = tokenLimit // Se o token existir, usa o limite específico
		}
	}

	// Increment the counter for the key
	count, err := rl.Store.Increment(key, rl.Expire)
	if err != nil {
		fmt.Println("Error incrementing key:", err)
		return false
	}

	// Apply appropriate limit
	if count > limit {
		fmt.Printf("Rate limit exceeded for '%s': '%s'\n", limitType, key)
		return false
	}

	return true
}
