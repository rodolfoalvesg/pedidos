package usecases

import (
	"context"
	"errors"
	"fmt"
	"order-api/app/domain/entities"

	"github.com/google/uuid"
)

func (u *Usecase) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*entities.Order, error) {
	const op = "UserUsecase.GetOrderByID"

	orderRDB, err := u.rdbRepository.GetOrder(ctx, orderID.String())
	if errors.Is(err, entities.ErrorOrderNotFoundInCache) {
		order, err := u.repository.GetOrderByID(ctx, orderID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if err := u.rdbRepository.SetOrder(ctx, order); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return order, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orderRDB, nil
}
