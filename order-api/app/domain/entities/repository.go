package entities

import (
	"context"

	"github.com/google/uuid"
)

type OrderRepository interface {
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*Order, error)
	GetAllOrders(ctx context.Context) ([]*Order, error)
	CreateOrder(ctx context.Context, order *Order) error
	UpdateOrder(ctx context.Context, order *Order) error
	DeleteOrder(ctx context.Context, orderID uuid.UUID) error
}

type OrderRDBRepository interface {
	SetOrder(ctx context.Context, order *Order) error
	GetOrder(ctx context.Context, orderID string) (*Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
}

type OrderElasticRepository interface {
	GetOrderIndex(ctx context.Context, filterOrder *OrderFilter) ([]*Order, error)
	IndexOrder(ctx context.Context, name string, order *Order) error
	DeleteOrder(ctx context.Context, publicID uuid.UUID) error
	UpdateOrder(ctx context.Context, order *Order) error
}

type UserClient interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) error
}
