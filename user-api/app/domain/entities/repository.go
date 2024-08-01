package entities

import (
	"context"

	"github.com/google/uuid"
)
//go:generate moq -out user_service_mock.go -pkg user ./user UserService
type UserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	UserExists(ctx context.Context, user *User) error
	GetAllUsers(ctx context.Context) ([]*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type UserRDBRepository interface {
	SetUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, userID string) error
}

type UserElasticRepository interface {
	GetUsersIndex(ctx context.Context, filter *UserFilter) ([]*User, error)
	UserIndex(ctx context.Context, name string, user *User) error
	DeleteUser(ctx context.Context, publicID uuid.UUID) error
	UpdateUser(ctx context.Context, user *User) error
}
