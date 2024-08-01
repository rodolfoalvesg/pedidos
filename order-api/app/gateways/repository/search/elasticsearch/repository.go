package elastic

import "github.com/elastic/go-elasticsearch/v8"

// UserElasticRepository is a struct to store elasticsearch client
type OrderElasticRepository struct {
	ES *elasticsearch.Client
}

// NewRepository is a function to create new elasticsearch repository
func NewRepository(es *elasticsearch.Client) *OrderElasticRepository {
	return &OrderElasticRepository{
		ES: es,
	}
}
