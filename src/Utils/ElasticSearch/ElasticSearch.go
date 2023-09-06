package ElasticSearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var ElasticConfig = elasticsearch.Config{
	Addresses: []string{
		"http://192.168.69.70:9331",
	},
}
