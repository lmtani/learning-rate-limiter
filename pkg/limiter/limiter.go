package limiter

import (
	"time"
)

type RateLimiterI interface {
	ShouldPass(key string) bool
}

type RateLimiterStore interface {
	Increment(key string, expire time.Duration) (int, error)
	Reset(key string) error
}

type RateLimiter struct {
	Limit  int
	Expire time.Duration
	Store  RateLimiterStore
}

func NewRateLimiter(limit int, expire time.Duration, store RateLimiterStore) *RateLimiter {
	return &RateLimiter{
		Limit:  limit,
		Expire: expire,
		Store:  store,
	}
}

func (rl *RateLimiter) ShouldPass(key string) bool {
	count, err := rl.Store.Increment(key, rl.Expire)
	if err != nil {
		return false
	}
	if count > rl.Limit {
		return false
	}
	return true
}
