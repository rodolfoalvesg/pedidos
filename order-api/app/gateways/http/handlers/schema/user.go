package schema

import (
	"order-api/app/domain/entities"
	"time"

	"github.com/google/uuid"
)

type OrderRequest struct {
	UserID          uuid.UUID `json:"user_id"`
	ItemDescription string    `json:"item_description"`
	ItemQuantity    int       `json:"item_quantity"`
	ItemPrice       float64   `json:"item_price"`
}

type UpdateOrderRequest struct {
	ItemDescription string  `json:"item_description"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
}

type OrderResponse struct {
	PublicID        uuid.UUID `json:"order_id"`
	UserID          uuid.UUID `json:"user_id"`
	ItemDescription string    `json:"item_description"`
	ItemQuantity    int       `json:"item_quantity"`
	ItemPrice       float64   `json:"item_price"`
	TotalValue      float64   `json:"total_value"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func MapOrderToResponse(order entities.Order) OrderResponse {
	return OrderResponse{
		PublicID:        order.PublicID,
		UserID:          order.UserID,
		ItemDescription: order.ItemDescription,
		ItemQuantity:    order.ItemQuantity,
		ItemPrice:       order.ItemPrice,
		TotalValue:      order.TotalValue,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
}

func MapOrdersToResponse(orders []*entities.Order) []entities.Order {
	var ordersResponse []entities.Order
	for _, order := range orders {
		ordersResponse = append(ordersResponse, *order)
	}

	return ordersResponse
}
