package redis

import (
	"order-api/app/domain/entities"

	"github.com/redis/go-redis/v9"
)

var _ entities.OrderRDBRepository = (*OrderRepository)(nil)

type OrderRepository struct {
	RDB *redis.Client
}

func NewRepository(rdb *redis.Client) *OrderRepository {
	return &OrderRepository{
		RDB: rdb,
	}
}
