package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"user-api/app/domain/entities"
)

func (r *UserRepository) GetUser(ctx context.Context, id string) (*entities.User, error) {
	const (
		op      = "UserRepository.GetUserRedis"
		service = "user-api:"
	)

	key := fmt.Sprintf("%s:%s", service, id)

	userFound, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, entities.ErrorUserNotFoundInCache)
	}

	var user *entities.User
	err = json.Unmarshal([]byte(userFound), &user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
