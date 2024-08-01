package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"order-api/app/domain/entities"
)

func (r *OrderRepository) SetOrder(ctx context.Context, order *entities.Order) error {
	const (
		op      = "OrderRepository.SetOrderRedis"
		service = "order-api:"
	)

	key := fmt.Sprintf("%s:%s", service, order.PublicID.String())

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = r.RDB.Set(ctx, key, orderJSON, 0).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
