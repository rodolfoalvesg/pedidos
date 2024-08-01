package usecases

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// GetAllUsers returns all users.
func (u *Usecase) GetAllUsers(ctx context.Context, filter *entities.UserFilter) ([]*entities.User, error) {
	const op = "UserUsecase.GetAllUsers"

	users, err := u.esRepository.GetUsersIndex(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}
