package ElasticSearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectElasticsearch() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"194.233.95.186:9331",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return es, nil
}
