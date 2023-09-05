package ElasticSearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
)

func GetElasticSearch(c *fiber.Ctx) error {
	// es, err := ElasticSearch.ConnectElasticsearch()
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Elasticsearch: %s", err)
	// }

	index := c.Query("index")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://194.233.95.186:9331",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal terhubung ke Elasticsearch")
	}

	response, err := es.Get(index, "counter_bookmark")

	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"counter_bookmark": 1,
	// 		},
	// 	},
	// }

	// req := esutil.NewJSONReader(query)

	// res, err := es.Search(
	// 	es.Search.WithIndex(index),
	// 	es.Search.WithBody(req),
	// 	es.Search.WithContext(context.Background()),
	// )

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString("Gagal melakukan pencarian data")
	// }

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-get-data-redis",
			"index":    index,
			"data":     response,
		},
	})
}
