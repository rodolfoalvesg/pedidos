package handlers

import (
	"log"
	"net/http"
	"order-api/app/domain/entities"
	"order-api/app/gateways/http/handlers/errors"
	"order-api/app/gateways/http/handlers/schema"
	"order-api/app/gateways/http/rest/requests"
	"order-api/app/gateways/http/rest/responses"
)

// UpdateOrder Update a order
// @Summary Update a order
// @Description Update a order
// @Tags orders
// @Accept json
//
// @Param oder-id path string true "Order ID"
// @Param Body body schema.UpdateOrderRequest true "Order Request Update"
//
// @Success 204 "No Content"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/order-api/orders/{order-id} [put]
func (h *Handler) UpdateOrder(r *http.Request) responses.Response {
	const op = "OrderHandler.UpdateOrder"

	orderID, err := requests.ParseUUID(r, "order-id")
	if err != nil {
		log.Printf("%s: failed to parse orderID: %v", op, err)

		return responses.BadRequest(err)
	}

	var req schema.UpdateOrderRequest
	if err := requests.DecodeBodyJSON(r, &req); err != nil {
		log.Printf("%s: failed to decode body JSON: %v", op, err)

		return responses.BadRequest(err)
	}

	order := &entities.Order{
		PublicID:        orderID,
		ItemDescription: req.ItemDescription,
		ItemQuantity:    req.ItemQuantity,
		ItemPrice:       req.ItemPrice,
	}

	if err := h.usecase.UpdateOrder(r.Context(), order); err != nil {
		log.Printf("%s: failed to update order: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.NoContent()
}
