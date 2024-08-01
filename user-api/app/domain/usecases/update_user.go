package usecases

import (
	"context"
	"fmt"
	"user-api/app/domain/entities"
)

// UpdateUser updates a user.
func (u *Usecase) UpdateUser(ctx context.Context, user *entities.User) error {
	const op = "UserUsecase.UpdateUser"

	userFound, err := u.repository.GetUserByID(ctx, user.PublicID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	user.CPF = userFound.CPF

	if userFound.Email != user.Email {
		if err := u.repository.UserExists(ctx, user); err != nil {
			return err
		}
	}

	if err := u.repository.UpdateUser(ctx, user); err != nil {
		return err
	}

	userRDB, _ := u.rdbRepository.GetUser(ctx, user.PublicID.String())
	if userRDB != nil {
		if err := u.rdbRepository.DeleteUser(ctx, user.PublicID.String()); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if err := u.esRepository.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
