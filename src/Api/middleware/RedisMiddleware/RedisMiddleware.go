package RedisMiddleware

import (
	"fmt"
	"go-learning/src/Utils/RedisClient"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/net/context"
)

var ctx = context.Background()

func Cache(c *fiber.Ctx) error {
	Client := RedisClient.Client
	stage := "production"
	if c.Get("stage") == "dev" {
		stage = "development"
	}

	expireTime := 30
	exStr := c.Query("ex")
	if exStr != "" {
		if ex, err := strconv.Atoi(exStr); err == nil {
			expireTime = ex
		}
	}

	skipCache := c.Query("nocache") != ""

	token := c.Get("Authorization")
	version := "V" + c.Get("V")
	if version == "V" {
		version = "V" + c.Get("v")
	}
	rediskey := ""
	if strings.HasPrefix(c.OriginalURL(), "/get/user") {
		rediskey = fmt.Sprintf("%s:%s:%s", c.OriginalURL(), version, stage)
	} else {
		rediskey = fmt.Sprintf("%s:%s:%s:%s", token, c.OriginalURL(), version, stage)
	}

	getFromRedis, err := Client.Get(ctx, rediskey).Result()

	if err == nil && !skipCache {
		c.Set("cached", "true")
		c.Set("Content-Type", "application/json")
		return c.Status(fiber.StatusOK).SendString(getFromRedis)
	}

	if skipCache {
		fmt.Println("delete", strings.Split(c.OriginalURL(), "?")[0])
		deleteFilter(ctx, fmt.Sprintf("*%s:%s*", token, strings.Split(c.OriginalURL(), "?")[0]))
	}

	// Override the response writer to capture the response
	capturedWriter := &responseCapturer{Response: c}
	c.Next()

	// Get the response data from the captured writer
	responseData := capturedWriter.ResponseData()
	log.Debug(responseData)

	if responseData.StatusCode == fiber.StatusOK && !skipCache {
		Client.Set(ctx, rediskey, responseData.Body, time.Duration(expireTime)*time.Second)
	}

	// Write the response data to the Fiber context and send it
	c.Status(responseData.StatusCode)
	c.Set("Content-Type", "application/json")
	c.SendString(responseData.Body)
	return nil
}

func deleteFilter(ctx context.Context, filter string) error {
	Client := RedisClient.Client
	keys, _, err := Client.Scan(ctx, 0, filter, 0).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		Client.Del(ctx, keys...)
	}

	return nil
}

type responseCapturer struct {
	Response *fiber.Ctx
}

func (r *responseCapturer) ResponseData() struct {
	StatusCode int
	Body       string
} {
	return struct {
		StatusCode int
		Body       string
	}{
		StatusCode: r.Response.Response().StatusCode(),
		Body:       string(r.Response.Response().Body()),
	}
}
