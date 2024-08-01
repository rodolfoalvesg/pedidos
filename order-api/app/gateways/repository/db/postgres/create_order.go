package postgres

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"
)

// CreateUser creates a new user.
func (r *OrderRepository) CreateOrder(ctx context.Context, order *entities.Order) error {
	const op = "OrderRepository.CreateOrder"

	if err := r.DB.Create(order).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
