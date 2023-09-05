package ElasticSearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var ElasticConfig = elasticsearch.Config{
	Addresses: []string{
		"http://194.233.95.186:9331",
	},
}
