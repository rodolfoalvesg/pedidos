package postgres

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"
)

func (r *OrderRepository) UpdateOrder(ctx context.Context, order *entities.Order) error {
	const op = "OrderRepository.UpdateOrder"

	// Update user with gorm not update all fields
	if err := r.DB.Model(&entities.Order{
		ItemDescription: order.ItemDescription,
		ItemQuantity:    order.ItemQuantity,
		ItemPrice:       order.ItemPrice,
		TotalValue:      order.TotalValue,
	}).Where("public_id = ?", order.PublicID).Updates(order).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
