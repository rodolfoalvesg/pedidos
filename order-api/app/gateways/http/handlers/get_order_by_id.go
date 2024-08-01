package handlers

import (
	"log"
	"net/http"
	"order-api/app/gateways/http/handlers/errors"
	"order-api/app/gateways/http/handlers/schema"
	"order-api/app/gateways/http/rest/requests"
	"order-api/app/gateways/http/rest/responses"
)

// GetOrderByID gets a order by ID
// @Summary Get a order by ID
// @Description Get a oder by ID
// @Tags orders
// @Accept json
//
// @Param order-id path string true "Order ID"
//
// @Success 200 {object} schema.OrderResponse "Success"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 404 {object} responses.NotFoundError "Not Found"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/order-api/orders/{order-id} [get]
func (h *Handler) GetOrderByID(r *http.Request) responses.Response {
	const op = "OrderHandler.GetOrderByID"

	orderID, err := requests.ParseUUID(r, "order-id")
	if err != nil {
		log.Printf("%s: failed to parse order ID: %v", op, err)

		return responses.BadRequest(err)
	}

	order, err := h.usecase.GetOrderByID(r.Context(), orderID)
	if err != nil {
		log.Printf("%s: failed to get order by email: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.OK(schema.MapOrderToResponse(*order))
}
