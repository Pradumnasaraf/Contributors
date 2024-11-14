package redis

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	limiter     *redis_rate.Limiter
)

func RedisInit() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URI"),
	})
	limiter = redis_rate.NewLimiter(redisClient)
}

func RedisClose() {
	err := redisClient.Close()
	if err != nil {
		log.Fatal("unable to close the redis connection.")
	}
}

func RateLimiter(clientIP string) error {
	ctx := context.Background()

	limitInt, _ := strconv.Atoi(os.Getenv("REDIS_RATE_LIMIT"))
	res, err := limiter.Allow(ctx, clientIP, redis_rate.PerHour(limitInt))
	if err != nil {
		log.Fatal("unable to connect to redis instance or check the limit of the incoming request.")
	}

	if res.Remaining == 0 {
		return errors.New("rate limit hit")
	}

	return nil
}
