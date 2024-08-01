package postgres

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// GetAllUsers returns all users.
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*entities.User, error) {
	const op = "UserRepository.GetAllUsers"

	users := []*entities.User{}
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return users, nil
}
