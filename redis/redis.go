package redis

import (
	"context"

	"github.com/FiberApps/common-library/logger"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect(uri string, user string, password string) {
	log := logger.New()
	Client = redis.NewClient(&redis.Options{
		Addr:     uri,
		Username: user,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Error("REDIS:: Error while connecting: %v", err)
		panic(err)
	}
	log.Info("REDIS:: Connected")
}
