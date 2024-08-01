package usecases

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"

	"github.com/google/uuid"
)

func (u *Usecase) DeleteOrder(ctx context.Context, orderID uuid.UUID) error {
	const op = "OrderUsecase.CreateOrder"

	var order *entities.Order

	orderFound, err := u.repository.GetOrderByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.repository.DeleteOrder(ctx, orderFound.PublicID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	orderRDB, _ := u.rdbRepository.GetOrder(ctx, orderFound.PublicID.String())
	if orderRDB != nil {
		if err := u.rdbRepository.DeleteOrder(ctx, order.PublicID.String()); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if err := u.esRepository.DeleteOrder(ctx, orderFound.PublicID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
