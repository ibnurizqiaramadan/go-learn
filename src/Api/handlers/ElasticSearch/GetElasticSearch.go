package ElasticSearch

import (
	"context"
	"encoding/json"

	"go-learning/src/Utils/ElasticSearch"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

func RetrieveElastic(c *fiber.Ctx) error {
	cfg := ElasticSearch.ElasticConfig

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal terhubung ke Elasticsearch")
	}

	request := esapi.GetRequest{
		Index:      "user_dev",  // Replace with your index name
		DocumentID: "user_doug", // Replace with the ID of the document you want to retrieve
	}

	// Execute the get request
	res, err := request.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error executing get request: %s", err)
	}

	// Decode the JSON response
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding JSON response: %s", err)
	}

	source, found := result["_source"]
	log.Debug(found)

	if found == false {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal melakukan pencarian data")
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success retrieve",
			"data":     source,
		},
	})
}
