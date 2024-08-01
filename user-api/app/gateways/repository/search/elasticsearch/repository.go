package elastic

import "github.com/elastic/go-elasticsearch/v8"

// UserElasticRepository is a struct to store elasticsearch client
type UserElasticRepository struct {
	ES *elasticsearch.Client
}

// NewRepository is a function to create new elasticsearch repository
func NewRepository(es *elasticsearch.Client) *UserElasticRepository {
	return &UserElasticRepository{
		ES: es,
	}
}
