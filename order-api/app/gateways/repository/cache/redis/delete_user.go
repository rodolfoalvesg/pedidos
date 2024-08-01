package redis

import (
	"context"
	"fmt"
)

func (r *OrderRepository) DeleteOrder(ctx context.Context, orderID string) error {
	const (
		op      = "OrderRDBRepository.DeleteOrder"
		service = "order-api:"
	)

	key := fmt.Sprintf("%s:%s", service, orderID)

	if err := r.RDB.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
