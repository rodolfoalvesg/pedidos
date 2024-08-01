package usecases

import (
	"context"
	"errors"
	"testing"
	"user-api/app/domain/entities"
	"user-api/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUsecase_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		input   entities.User
		repoDB  entities.UserRepository
		repoRDB entities.UserRDBRepository
		repoES  entities.UserElasticRepository
	}{
		{
			name: "success when creating a new user",
			input: entities.User{
				Name:  "John Doe",
				CPF:   "12345678901",
				Email: "jh@t.com",
				Phone: "1933330000",
			},
			repoDB: &mocks.UserRepositoryMock{
				UserExistsFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
				CreateUserFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
			},
			repoRDB: &mocks.UserRDBRepositoryMock{
				SetUserFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
			},
			repoES: &mocks.UserElasticRepositoryMock{
				UserIndexFunc: func(ctx context.Context, name string, user *entities.User) error {
					return nil
				},
			},
		},
		{
			name: "error when user already exists",
			input: entities.User{
				Name:  "John Doe",
				CPF:   "12345678901",
				Email: "jh@t.com",
				Phone: "1933330000",
			},
			repoDB: &mocks.UserRepositoryMock{
				UserExistsFunc: func(ctx context.Context, user *entities.User) error {
					return entities.ErrUserAlreadyExists
				},
			},
		},
		{
			name: "error creating a new user",
			input: entities.User{
				Name:  "John Doe",
				CPF:   "12345678901",
				Email: "jh@t.com",
				Phone: "1933330000",
			},
			repoDB: &mocks.UserRepositoryMock{
				UserExistsFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
				CreateUserFunc: func(ctx context.Context, user *entities.User) error {
					return errors.New("Internal Server Error")
				},
			},
		},
		{
			name: "error when creating a new user in redis",
			input: entities.User{
				Name:  "John Doe",
				CPF:   "12345678901",
				Email: "jh@t.com",
				Phone: "1933330000",
			},
			repoDB: &mocks.UserRepositoryMock{
				UserExistsFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
				CreateUserFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
			},
			repoRDB: &mocks.UserRDBRepositoryMock{
				SetUserFunc: func(ctx context.Context, user *entities.User) error {
					return errors.New("Internal Server Error")
				},
			},
		},
		{
			name: "error when creating a new user in elasticsearch",
			input: entities.User{
				Name:  "John Doe",
				CPF:   "12345678901",
				Email: "jh@t.com",
				Phone: "1933330000",
			},
			repoDB: &mocks.UserRepositoryMock{
				UserExistsFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
				CreateUserFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
			},
			repoRDB: &mocks.UserRDBRepositoryMock{
				SetUserFunc: func(ctx context.Context, user *entities.User) error {
					return nil
				},
			},
			repoES: &mocks.UserElasticRepositoryMock{
				UserIndexFunc: func(ctx context.Context, name string, user *entities.User) error {
					return errors.New("Internal Server Error")
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := NewUsecase(tt.repoDB, tt.repoRDB, tt.repoES)

			err := u.CreateUser(context.Background(), &tt.input)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}

}
