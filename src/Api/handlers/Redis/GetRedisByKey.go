package Redis

import (
	"context"
	"go-learning/src/Utils/RedisClient"
	"go-learning/src/Utils/Validation"

	"github.com/gofiber/fiber/v2"
)

type GetRedisKey struct {
	Key string `json:"key"`
}

func GetRedisByKey(c *fiber.Ctx) error {
	key := GetRedisKey{}
	ctx := context.Background()
	Client := RedisClient.Client

	errors, isValid := Validation.ValidateInput(c, key)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"statusCode": fiber.StatusBadRequest, "messages": "Invalid Input", "errors": errors})
	}

	data, errRedis := Client.Get(ctx, key.Key).Result()
	if errRedis != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "Failed get data from redis: " + errRedis.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-get-data-redis",
			"key":      key.Key,
			"data":     data,
		},
	})
}
