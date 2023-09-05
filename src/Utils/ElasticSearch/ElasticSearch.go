package ElasticSearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var cfg = elasticsearch.Config{
	Addresses: []string{
		"http://194.233.95.186:9331",
	},
}

func ConnectElasticsearch() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://194.233.95.186:9331",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return es, nil
}
