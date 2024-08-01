package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"user-api/app/domain/entities"
)

func (r *UserRepository) SetUser(ctx context.Context, user *entities.User) error {
	const (
		op      = "UserRepository.SetUserRedis"
		service = "user-api:"
	)

	key := fmt.Sprintf("%s:%s", service, user.PublicID.String())

	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = r.RDB.Set(ctx, key, userJSON, 0).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
