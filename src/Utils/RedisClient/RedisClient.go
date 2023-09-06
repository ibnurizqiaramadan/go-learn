package RedisClient

import (
	"go-learning/src/Utils/Functions"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       Functions.StoI(os.Getenv("REDIS_DB")),
	})
}
