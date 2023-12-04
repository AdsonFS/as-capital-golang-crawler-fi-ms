package redis_repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var clientRedis *redis.Client

func Get(ctx context.Context, key string) map[string]string {
	return getClient().HGetAll(ctx, key).Val()
}

func Save(ctx context.Context, key string, values map[string]float64) error {
	for k, v := range values {
		err := getClient().HSet(ctx, key, k, v).Err()
		if err != nil {
			return err
		}
	}
	err := getClient().Expire(ctx, key, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func getClient() *redis.Client {
	if clientRedis == nil {
		clientRedis = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	}
	return clientRedis
}
