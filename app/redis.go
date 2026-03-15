package main

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func initRedis() {

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: redisHost,
	})
}

func incrementVisits() int64 {
	val, err := rdb.Incr(ctx, "visits").Result()
	if err != nil {
		return -1
	}
	return val
}
