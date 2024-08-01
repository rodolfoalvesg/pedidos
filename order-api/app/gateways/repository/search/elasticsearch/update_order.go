package elastic

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"
)

func (r *OrderElasticRepository) UpdateOrder(ctx context.Context, order *entities.Order) error {
	const op = "OrderElasticRepository.UpdateOrder"

	if err := r.IndexOrder(ctx, "orders", order); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
