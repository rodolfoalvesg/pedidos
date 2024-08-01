package postgres

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"

	"github.com/google/uuid"
)

// GetUserByID fetches a user by its ID
func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	const op = "UserRepository.GetUserByID"

	user := &entities.User{}
	if err := r.DB.Where("public_id = ?", id).First(user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("%s: %w", op, entities.ErrorUserNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
