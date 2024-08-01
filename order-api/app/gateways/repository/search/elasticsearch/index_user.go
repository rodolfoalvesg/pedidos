package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"order-api/app/domain/entities"
	"strings"
)

func (r *OrderElasticRepository) IndexOrder(ctx context.Context, name string, order *entities.Order) error {
	const op = "UserElasticRepository.IndexUser"

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	resp, err := r.ES.Index(
		name,
		strings.NewReader(string(orderJSON)),
		r.ES.Index.WithDocumentID(order.PublicID.String()),
		r.ES.Index.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer resp.Body.Close()

	return nil

}
