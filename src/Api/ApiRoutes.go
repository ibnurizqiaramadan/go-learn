package ApiRoutes

import (
	"go-learning/src/Api/handlers/BasicUsages"
	"go-learning/src/Api/handlers/Billings"
	"go-learning/src/Api/handlers/ElasticSearch"
	"go-learning/src/Api/handlers/HowToGetQuery"
	"go-learning/src/Api/handlers/Index"
	QueryMutation "go-learning/src/Api/handlers/Query"
	"go-learning/src/Api/handlers/RabbitMQ"
	"go-learning/src/Api/handlers/Redis"
	"go-learning/src/Api/handlers/Users"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(http *fiber.App) {
	http.Get("/", Index.Index)

	// basic ussages
	http.Get("/basic/:name<int>", BasicUsages.GetDatas)

	// Users Routes Hasura Graphql
	http.Get("/users", Users.AddUsers)
	http.Get("/users/:name", Users.AddUsers)
	http.Get("/get-admin", HowToGetQuery.ExampleGetUsingAdmin)
	http.Get("/get-user", HowToGetQuery.ExampleGetUsingUser)
	http.Get("/get-pagination", HowToGetQuery.ExampleGetPagination)
	http.Get("/get-where", HowToGetQuery.ExampleGetWhere)
	http.Post("/mutation", QueryMutation.MutationUsers)
	http.Post("/mutation-update", QueryMutation.MutationUpdateUsers)

	// Billings Routes Stripe
	BillingGroup := http.Group("billing", func(c *fiber.Ctx) error { return c.Next() })
	BillingGroup.Post("/create-customer", Billings.AddCustomer)
	BillingGroup.Get("/get-customer", Billings.GetCustomer)

	// Users Routes Mysql
	http.Get("/test", Users.TestConnection)
	http.Get("/users/:id", Users.GetUserDetails)
	http.Post("/login", Users.Login)
	http.Post("/user/add", Users.AddUserMysql)

	// Redis Routes
	http.Get("/redis", Redis.GetRedisByKey)

	// ElasticSearch Routes
	http.Get("/elasticsearch", ElasticSearch.GetElasticSearch)
	http.Get("/es-retrieve", ElasticSearch.RetrieveElastic)

	// UPLOAD FILE
	http.Post("/upload", BasicUsages.UploadFile)

	// RabbitMQ
	http.Get("/rabbitmq", RabbitMQ.SendRebbitMQ)
}
