package usecases

import (
	"context"
	"user-api/app/domain/entities"

	"github.com/google/uuid"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetAllUsers(ctx context.Context, filter *entities.UserFilter) ([]*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type Usecase struct {
	repository    entities.UserRepository
	rdbRepository entities.UserRDBRepository
	esRepository  entities.UserElasticRepository
}

func NewUsecase(
	db entities.UserRepository,
	rdb entities.UserRDBRepository,
	es entities.UserElasticRepository,
) *Usecase {
	return &Usecase{
		repository:    db,
		rdbRepository: rdb,
		esRepository:  es,
	}
}
