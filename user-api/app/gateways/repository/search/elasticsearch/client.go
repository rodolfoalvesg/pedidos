package elastic

import (
	"user-api/config"

	"github.com/elastic/go-elasticsearch/v8"
)

// NewClient is a function to create new elasticsearch client
func NewClient(cfg *config.Config) (*elasticsearch.Client, error) {
	cfgES := elasticsearch.Config{
		Addresses: []string{
			cfg.ES.Host,
		},
		Username: cfg.ES.Username,
		Password: cfg.ES.Password,
	}

	es, err := elasticsearch.NewClient(cfgES)
	if err != nil {
		return nil, err
	}

	return es, nil

}
