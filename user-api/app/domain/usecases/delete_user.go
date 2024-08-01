package usecases

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"

	"github.com/google/uuid"
)

// DeleteUser deletes a user.
func (u *Usecase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	const op = "UserUsecase.DeleteUser"

	var user *entities.User

	userFound, err := u.repository.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := u.repository.DeleteUser(ctx, userFound.PublicID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	user, _ = u.rdbRepository.GetUser(ctx, userFound.PublicID.String())
	if user != nil {
		if err := u.rdbRepository.DeleteUser(ctx, user.PublicID.String()); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if err := u.esRepository.DeleteUser(ctx, userFound.PublicID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
