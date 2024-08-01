package redis

import (
	"user-api/app/domain/entities"

	"github.com/redis/go-redis/v9"
)

var _ entities.UserRDBRepository = (*UserRepository)(nil)

type UserRepository struct {
	RDB *redis.Client
}

func NewRepository(rdb *redis.Client) *UserRepository {
	return &UserRepository{
		RDB: rdb,
	}
}
