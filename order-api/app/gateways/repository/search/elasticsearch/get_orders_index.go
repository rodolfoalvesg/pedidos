package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"order-api/app/domain/entities"
	"strings"
)

// GetUserIndex is a function to get user index in elasticsearch
func (r *OrderElasticRepository) GetOrderIndex(ctx context.Context, filter *entities.OrderFilter) ([]*entities.Order, error) {
	const op = "OrderElasticRepository.GetOrdersIndex"

	var buf strings.Builder
	buf.WriteString(`{"query": {`)

	queryParts := []string{}

	if filter.UserID != "" {
		queryParts = append(queryParts, fmt.Sprintf(`{"match": {"user_id": "%s"}}`, filter.UserID))
	}

	if filter.PublicID != "" {
		queryParts = append(queryParts, fmt.Sprintf(`{"match": {"public_id": "%s"}}`, filter.PublicID))
	}

	if filter.Description != "" {
		queryParts = append(queryParts, fmt.Sprintf(`{"match": {"item_description": "%s"}}`, filter.Description))
	}

	if len(queryParts) > 0 {
		buf.WriteString(fmt.Sprintf(`"bool": {"must": [%s]}`, strings.Join(queryParts, ",")))
	} else {
		buf.WriteString(`"match_all": {}`)
	}

	buf.WriteString(`}}`)

	res, err := r.ES.Search(
		r.ES.Search.WithContext(ctx),
		r.ES.Search.WithIndex("orders"),
		r.ES.Search.WithBody(strings.NewReader(buf.String())),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer res.Body.Close()

	// Process the response
	var rMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&rMap); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	hits := rMap["hits"].(map[string]interface{})["hits"].([]interface{})
	users := make([]*entities.Order, 0, len(hits))
	for _, hit := range hits {
		var user entities.Order
		doc := hit.(map[string]interface{})["_source"]
		userData, _ := json.Marshal(doc)
		if err := json.Unmarshal(userData, &user); err != nil {
			return nil, fmt.Errorf("error unmarshalling user: %w", err)
		}
		users = append(users, &user)
	}

	return users, nil
}
