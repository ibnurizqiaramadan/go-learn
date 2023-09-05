package ApiRoutes

import (
	"go-learning/src/Api/handlers/BasicUsages"
	"go-learning/src/Api/handlers/Billings"
	"go-learning/src/Api/handlers/ElasticSearch"
	"go-learning/src/Api/handlers/HowToGetQuery"
	"go-learning/src/Api/handlers/Index"
	QueryMutation "go-learning/src/Api/handlers/Query"
	"go-learning/src/Api/handlers/Redis"
	"go-learning/src/Api/handlers/Users"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(http *fiber.App) {
	http.Get("/", func(c *fiber.Ctx) error { return Index.Index(c) })

	// basic ussages
	http.Get("/basic/:name<int>", func(c *fiber.Ctx) error { return BasicUsages.GetDatas(c) })

	// Users Routes Hasura Graphql
	http.Get("/users", func(c *fiber.Ctx) error { return Users.AddUsers(c) })
	http.Get("/get-admin", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetUsingAdmin(c) })
	http.Get("/get-user", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetUsingUser(c) })
	http.Get("/get-pagination", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetPagination(c) })
	http.Get("/get-where", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetWhere(c) })
	http.Post("/mutation", func(c *fiber.Ctx) error { return QueryMutation.MutationUsers(c) })
	http.Post("/mutation-update", func(c *fiber.Ctx) error { return QueryMutation.MutationUpdateUsers(c) })

	// Billings Routes Stripe
	BillingGroup := http.Group("billing", func(c *fiber.Ctx) error { return c.Next() })
	BillingGroup.Post("/create-customer", func(c *fiber.Ctx) error { return Billings.AddCustomer(c) })
	BillingGroup.Get("/get-customer", func(c *fiber.Ctx) error { return Billings.GetCustomer(c) })

	// Users Routes Mysql
	http.Get("/test", func(c *fiber.Ctx) error { return Users.TestConnection(c) })
	http.Get("/users/:id", func(c *fiber.Ctx) error { return Users.GetUserDetails(c) })
	http.Get("/token", func(c *fiber.Ctx) error { return Users.GenerateToken(c) })

	// Redis Routes
	http.Get("/redis", func(c *fiber.Ctx) error { return Redis.GetRedisByKey(c) })

	// ElasticSearch Routes
	http.Get("/elasticsearch", func(c *fiber.Ctx) error { return ElasticSearch.GetElasticSearch(c) })
	http.Get("/es-retrieve", func(c *fiber.Ctx) error { return ElasticSearch.RetrieveElastic(c) })
}
