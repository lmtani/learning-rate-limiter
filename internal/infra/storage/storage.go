package storage

import (
	"context"
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
func (r *RedisStorage) Increment(key string, expire time.Duration) (int, error) {
	pipe := r.Client.TxPipeline()
	incr := pipe.Incr(context.TODO(), key)
	pipe.Expire(context.TODO(), key, expire)
	_, err := pipe.Exec(r.Client.Context())
	if err != nil {
		return 0, err
	}
	// Retorna o valor incrementado
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
