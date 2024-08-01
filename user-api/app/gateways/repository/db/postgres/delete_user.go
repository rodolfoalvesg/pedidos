package postgres

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"

	"github.com/google/uuid"
)

// DeleteUser deletes a user by its ID
func (r *UserRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	const op = "UserRepository.DeleteUser"

	if err := r.DB.Where("public_id = ?", userID).Delete(&entities.User{}).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
