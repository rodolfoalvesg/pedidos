package redis

import (
	"context"
	"fmt"
)

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	const (
		op      = "UserRDBRepository.DeleteUser"
		service = "user-api:"
	)

	key := fmt.Sprintf("%s:%s", service, userID)

	if err := r.RDB.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
