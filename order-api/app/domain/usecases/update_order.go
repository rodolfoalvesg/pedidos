package usecases

import (
	"context"
	"fmt"
	"order-api/app/domain/entities"
)

func (u *Usecase) UpdateOrder(ctx context.Context, order *entities.Order) error {
	const op = "OrderUsecase.UpdateOrder"

	_, err := u.repository.GetOrderByID(ctx, order.PublicID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.repository.UpdateOrder(ctx, order); err != nil {
		return err
	}

	orderRDB, _ := u.rdbRepository.GetOrder(ctx, order.PublicID.String())
	if orderRDB != nil {
		if err := u.rdbRepository.DeleteOrder(ctx, order.PublicID.String()); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if err := u.esRepository.UpdateOrder(ctx, order); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
