package middleware

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

func RateLimiter(clientIP string) error {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	limiter := redis_rate.NewLimiter(rdb)
	limitInt, _ := strconv.Atoi(os.Getenv("REDIS_RATE_LIMIT"))
	res, err := limiter.Allow(ctx, clientIP, redis_rate.PerHour(limitInt))
	if err != nil {
		panic(err)
	}

	if res.Remaining == 0 {
		return errors.New("rate limit hit")
	}

	return nil
}
