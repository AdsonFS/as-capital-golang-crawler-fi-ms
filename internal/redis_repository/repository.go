package redis_repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisRepository *RedisRepository

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository() *RedisRepository {
	if redisRepository == nil {
		redisRepository = &RedisRepository{}
	}
	if redisRepository.client == nil {
		redisRepository.client = redis.NewClient(&redis.Options{
			Addr: ":6379",
		})
	}
	return redisRepository
}

func (repository *RedisRepository) Get(ctx context.Context, key string) map[string]string {
	return repository.client.HGetAll(ctx, key).Val()
}

func (repository *RedisRepository) Save(ctx context.Context, key string, values map[string]float64) error {
	for k, v := range values {
		err := repository.client.HSet(ctx, key, k, v).Err()
		if err != nil {
			return err
		}
	}
	err := repository.client.Expire(ctx, key, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}
