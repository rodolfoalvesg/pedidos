package postgres

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// UserExists checks if a user exists by email or cpf.
func (r *UserRepository) UserExists(ctx context.Context, user *entities.User) error {
	const op = "UserRepository.UserExists"

	if err := r.DB.Where("email = ?", user.Email).Or("cpf = ?", user.CPF).First(&entities.User{}).Error; err != nil {
		if err.Error() == "record not found" {
			return nil
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return fmt.Errorf("%s: %w", op, entities.ErrUserAlreadyExists)
}
