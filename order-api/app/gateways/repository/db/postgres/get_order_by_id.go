package postgres

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"

	"github.com/google/uuid"
)

func (r *OrderRepository) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*entities.Order, error) {
	const op = "OrderRepository.GetOrderByID"

	order := &entities.Order{}
	if err := r.DB.Where("public_id = ?", orderID).First(order).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("%s: %w", op, entities.ErrorOrderNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}
