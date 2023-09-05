package ElasticSearch

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/gofiber/fiber/v2"

	"go-learning/src/Utils/ElasticSearch"
)

func GetElasticSearch(c *fiber.Ctx) error {
	index := c.Query("index")
	cfg := ElasticSearch.ElasticConfig

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal terhubung ke Elasticsearch")
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"counter_bookmark": 1,
			},
		},
	}
	req := esutil.NewJSONReader(query)

	res, err := es.Search(
		es.Search.WithIndex(index),
		es.Search.WithBody(req),
		es.Search.WithContext(context.Background()),
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal melakukan pencarian data")
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-get-data-redis",
			"index":    index,
			"data":     res,
		},
	})
}
