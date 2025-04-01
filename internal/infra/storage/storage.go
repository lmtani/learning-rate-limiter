package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	Client *redis.Client
}

func NewRedisStorage(addr, password string, db int) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{Client: client}
}

// Increment increments the value of the key in Redis and sets an expiration time.
// The expiration time is set to the value of 'expire' and is updated with each increment.
// The method returns the incremented value and an error if any occurs.
func (r *RedisStorage) Increment(key string, expire time.Duration, limit int) (int, error) {
	ctx := context.TODO()

	// Get current value first to check limit
	val, err := r.Client.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	// Check if already at or over limit
	if val >= limit {
		return 0, fmt.Errorf("limit of %d reached for key: %s", val, key)
	}

	// Increment the counter
	incr := r.Client.Incr(ctx, key)
	if err := incr.Err(); err != nil {
		return 0, err
	}

	// Reset expiration time
	if err := r.Client.Expire(ctx, key, expire).Err(); err != nil {
		return 0, err
	}

	return int(incr.Val()), nil
}

// Reset removes the key from Redis.
func (r *RedisStorage) Reset(key string) error {
	err := r.Client.Del(context.TODO(), key).Err()
	if err != nil {
		return err
	}
	return nil
}
