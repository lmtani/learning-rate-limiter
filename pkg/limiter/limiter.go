package limiter

import (
	"time"

	"github.com/lmtani/learning-rate-limiter/internal/entity"
)

type RateLimiterI interface {
	ShallPass(key string) bool
}

type RateLimiterStore interface {
	Increment(key string, expire time.Duration, limit int) (int, error)
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

func (rl *RateLimiter) ShallPass(key string, limitType string) bool {
	if limitType != "api_key" && limitType != "ip" {
		return false
	}

	limit := rl.Limit // Na falta de api_key, limite padrão é mesmo utilizado para IP
	if limitType == "api_key" {
		if tokenLimit, exists := rl.TokenMap[key]; exists {
			limit = tokenLimit // Se o token existir, usa o limite específico
		}
	}

	// Increment the counter for the key
	_, err := rl.Store.Increment(key, rl.Expire, limit)
	if err != nil {
		return false
	}

	return true
}
