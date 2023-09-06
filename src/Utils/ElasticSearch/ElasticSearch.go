package ElasticSearch

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var ElasticConfig = elasticsearch.Config{
	Addresses: []string{
		os.Getenv("ELASTICSEARCH_ADDRESS"),
	},
}
