package caches

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dmzsz/duozhuayu/internal/configs"
	"github.com/redis/go-redis/v9"
)

type Cmdable interface {
	Set(key string, value interface{}, keyExpiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
}

var _ Cmdable = (*RedisCache)(nil)

type RedisCache struct {
	addr       string
	db         int
	password   string
	expiration time.Duration
	client     *redis.Client
}

var redisCache *RedisCache

// var redisConfig configs.REDIS = configs.AppConfig.DatabaseConfig.REDIS

func NewRedisCache(options *redis.Options, expiration time.Duration) *RedisCache {
	redisConfig := configs.AppConfig.DatabaseConfig.REDIS
	if options.Addr == "" {
		options.Addr = redisConfig.Env.Addr
	}
	if options.DB == 0 {
		options.DB = redisConfig.Env.DB
	}
	if options.Password == "" {
		options.Password = redisConfig.Env.Password
	}
	if expiration == 0 {
		expiration = redisConfig.Conn.Expiration
	}

	redisCache = &RedisCache{
		addr:       options.Addr,
		db:         options.DB,
		password:   options.Password,
		expiration: expiration,
		client: redis.NewClient(&redis.Options{
			Addr:     options.Addr,
			DB:       options.DB,
			Password: options.Password,
		}),
	}

	return redisCache
}

func GetRedisCache() *RedisCache {
	return redisCache
}

func (cache *RedisCache) Set(key string, value interface{}, keyExpiration time.Duration) (err error) {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cache.expiration*time.Second)
	defer cancel()

	err = cache.client.Set(ctx, key, json, time.Minute*cache.expiration).Err()

	if keyExpiration != 0 {
		err = cache.client.Expire(ctx, key, keyExpiration).Err()
	}
	return err
}

func (cache *RedisCache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cache.expiration*time.Second)
	defer cancel()

	return cache.client.Get(ctx, key).Result()
}

func (cache *RedisCache) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), cache.expiration*time.Second)
	defer cancel()

	return cache.client.Del(ctx, key).Err()
}
