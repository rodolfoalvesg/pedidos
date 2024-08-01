package postgres

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"

	"github.com/google/uuid"
)

func (r *OrderRepository) DeleteOrder(ctx context.Context, orderID uuid.UUID) error {
	const op = "OrderRepository.DeleteOrder"

	if err := r.DB.Where("public_id = ?", orderID).Delete(&entities.Order{}).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
