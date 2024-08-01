package elastic

import (
	"context"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func CreateIndex(ctx context.Context, name string, es *elasticsearch.Client) error {
	const op = "OrderElasticRepository.CreateOrderIndex"

	resp, err := es.Indices.Create(
		name,
		es.Indices.Create.WithContext(ctx),
		es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	return nil
}

func IndexExists(ctx context.Context, name string, es *elasticsearch.Client) (bool, error) {
	const op = "OrderElasticRepository.IndexExists"

	resp, err := es.Indices.Exists(
		[]string{name},
		es.Indices.Exists.WithContext(ctx),
	)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}
