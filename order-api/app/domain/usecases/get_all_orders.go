package usecases

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"
)

func (u *Usecase) GetAllOrders(ctx context.Context, filter *entities.OrderFilter) ([]*entities.Order, error) {
	const op = "OrderUsecase.GetAllOrders"

	orders, err := u.esRepository.GetOrderIndex(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}
