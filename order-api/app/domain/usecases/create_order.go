package usecases

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"
)

// CreateUser creates a new user.
func (u *Usecase) CreateOrder(ctx context.Context, order *entities.Order) error {
	const op = "OrderUsecase.CreateOrder"

	err := u.client.GetUserByID(ctx, order.UserID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.repository.CreateOrder(ctx, order); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.rdbRepository.SetOrder(ctx, order); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.esRepository.IndexOrder(ctx, "orders", order); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
