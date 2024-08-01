package postgres

import (
	"user-api/app/domain/entities"

	"gorm.io/gorm"
)

var _ entities.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
