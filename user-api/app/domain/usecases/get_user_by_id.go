package usecases

import (
	"context"
	"errors"
	"fmt"
	"user-api/app/domain/entities"

	"github.com/google/uuid"
)

// GetUserByID returns a user by its ID.
func (u *Usecase) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	const op = "UserUsecase.GetUserByID"

	userRDB, err := u.rdbRepository.GetUser(ctx, id.String())
	if errors.Is(err, entities.ErrorUserNotFoundInCache) {
		user, err := u.repository.GetUserByID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if err := u.rdbRepository.SetUser(ctx, user); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return user, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return userRDB, nil
}
