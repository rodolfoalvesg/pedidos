package elastic

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// UpdateUser is a function to update user in elasticsearch
func (r *UserElasticRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	const op = "UserElasticRepository.UpdateUser"

	if err := r.UserIndex(ctx, "users", user); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
