package postgres

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// CreateUser creates a new user.
func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	const op = "UserRepository.CreateUser"

	if err := r.DB.Create(user).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
