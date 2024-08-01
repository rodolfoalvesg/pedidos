package postgres

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// UpdateUser updates a user in the database
func (r *UserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	const op = "UserRepository.UpdateUser"

	// Update user with gorm not update all fields
	if err := r.DB.Model(&entities.User{
		Name:  user.Name,
		CPF:   user.CPF,
		Email: user.Email,
		Phone: user.Phone,
	}).Where("public_id = ?", user.PublicID).Updates(user).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
