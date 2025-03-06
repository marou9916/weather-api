package configs

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost: 6379",
		Password: "",
		DB:       0,
	})

	err := testRedisConnection()
	if err != nil {
		panic("Erreur lors de la connexion à Redis : " + err.Error())
	}
}

func testRedisConnection() error {
	var err error

	ctx, cancel := context.WithTimeout(Ctx, 5*time.Second)
	defer cancel()

	_, err = RedisClient.Ping(ctx).Result()

	return err
}
