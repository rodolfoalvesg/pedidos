package postgres

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"
)

func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]*entities.Order, error) {
	const op = "OrderRepository.GetAllOrders"

	orders := []*entities.Order{}
	if err := r.DB.Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return orders, nil
}
