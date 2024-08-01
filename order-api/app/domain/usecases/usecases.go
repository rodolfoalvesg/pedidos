package usecases

import (
	"context"
	"order-api/app/domain/entities"
	"order-api/app/gateways/http/client"

	"github.com/google/uuid"
)

type OrderUsecase interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (*entities.Order, error)
	GetAllOrders(ctx context.Context, filter *entities.OrderFilter) ([]*entities.Order, error)
	CreateOrder(ctx context.Context, order *entities.Order) error
	UpdateOrder(ctx context.Context, order *entities.Order) error
	DeleteOrder(ctx context.Context, orderID uuid.UUID) error
}

type Usecase struct {
	repository    entities.OrderRepository
	rdbRepository entities.OrderRDBRepository
	esRepository  entities.OrderElasticRepository
	client        *client.Client
}

func NewUsecase(
	db entities.OrderRepository,
	rdb entities.OrderRDBRepository,
	es entities.OrderElasticRepository,
	client *client.Client,
) *Usecase {
	return &Usecase{
		repository:    db,
		rdbRepository: rdb,
		esRepository:  es,
		client:        client,
	}
}
