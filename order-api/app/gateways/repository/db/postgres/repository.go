package postgres

import (
	"order-api/app/domain/entities"

	"gorm.io/gorm"
)

var _ entities.OrderRepository = (*OrderRepository)(nil)

type OrderRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}
