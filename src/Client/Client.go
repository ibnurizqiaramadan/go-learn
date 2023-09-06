package Client

import (
	"fmt"

	ApiRoutes "go-learning/src/Api"
	Middleware "go-learning/src/Api/middleware"
	"go-learning/src/Interfaces"
	"go-learning/src/Utils/RedisClient"
	"go-learning/src/Utils/StripeClient"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunServer(params Interfaces.SystemInterface) {
	app := fiber.New(fiber.Config{
		Prefork:       params.AppPrefork,
		CaseSensitive: params.AppCaseSensitive,
		StrictRouting: params.AppStrictRouting,
		ServerHeader:  params.ServerHeader,
		AppName:       params.AppName,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		BodyLimit:     params.AppBodyLimit * 1024 * 1024,
	})

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: params.CorsAllowOrigin,
		AllowMethods: params.CorsAllowMethod,
		AllowHeaders: params.CorsAllowHeader,
	}))

	app.Use(logger.New())                                                // logging
	app.Use(func(c *fiber.Ctx) error { return Middleware.CheckAuth(c) }) // auth middleware check

	StripeClient.InitStripe()
	RedisClient.InitRedis()

	ApiRoutes.InitRoutes(app)
	app.Get("/metrics", monitor.New())

	app.Listen(params.AppListenHost + ":" + fmt.Sprint(params.AppListenPort))
}
