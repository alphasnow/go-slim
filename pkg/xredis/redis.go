package xredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedis(cfg *Config) *redis.Client {
	// https://redis.uptrace.dev/guide/
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.GetAddr(),
		Password: cfg.Pass, // no password set
		DB:       cfg.DB,   // use default DB
		// PoolSize: 1000, //
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	return rdb
}
