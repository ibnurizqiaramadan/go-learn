package RedisClient

import (
	"github.com/redis/go-redis/v9"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     "194.233.95.186:6381",
	Password: "eP5gtsmn8SSSUfZQBkJIcaj0pcy8t+c4XPuiPik8gxMOan6XoTCLQuDXV8g+nLRIYpuAYdgywu9gJB+X",
	DB:       0,
})
