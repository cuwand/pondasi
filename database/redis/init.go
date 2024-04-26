package redis

import (
	"context"
	"fmt"
	"github.com/cuwand/pondasi/logger"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type Redis struct {
	client *redis.Client
	logger logger.Logger
}

var redisClient Redis

func InitConnection(redisDB, redisHost, redisPort, redisPassword string, logger logger.Logger) Redis {
	db := 0

	parseRedisDb, err := strconv.ParseInt(redisDB, 10, 32)

	if err == nil {
		db = int(parseRedisDb)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", redisHost, redisPort),
		Password: redisPassword,
		DB:       db,
	})

	if client.Ping(context.Background()).Err() != nil {
		logger.Error("cannot connect redis")
		panic("cannot connect redis")
	}

	logger.Info("Redis Connected")

	redisClient = Redis{
		client: client,
		logger: logger,
	}

	return redisClient
}

func (r Redis) GetRedisClient() *redis.Client {
	return r.client
}

func (r Redis) GetRedisLogger() logger.Logger {
	return r.logger
}

func GetClient() Collections {
	return redisClient
}
