package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"order-api/app/domain/entities"
)

func (r *OrderRepository) GetOrder(ctx context.Context, orderID string) (*entities.Order, error) {
	const (
		op      = "OrderRepository.GetOrderRedis"
		service = "order-api:"
	)

	key := fmt.Sprintf("%s:%s", service, orderID)

	orderFound, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, entities.ErrorOrderNotFoundInCache)
	}

	var order *entities.Order
	err = json.Unmarshal([]byte(orderFound), &order)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}
