package usecases

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// CreateUser creates a new user.
func (u *Usecase) CreateUser(ctx context.Context, user *entities.User) error {
	const op = "UserUsecase.CreateUser"

	if err := u.repository.UserExists(ctx, user); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.repository.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.rdbRepository.SetUser(ctx, user); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.esRepository.UserIndex(ctx, "users", user); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
