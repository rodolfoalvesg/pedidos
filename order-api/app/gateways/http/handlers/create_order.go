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

// CreateOrder creates a new order.
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept json
//
// @Param Body body schema.OrderRequest true "Order Request"
//
// @Success 201 "Created"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/order-api/orders [post]
func (h *Handler) CreateOrder(r *http.Request) responses.Response {
	const op = "OrderHandler.CreateOrder"

	var req schema.OrderRequest
	if err := requests.DecodeBodyJSON(r, &req); err != nil {
		log.Printf("%s: failed to decode body JSON: %v", op, err)

		return responses.BadRequest(err)
	}

	user, err := entities.NewOrder(req.UserID, req.ItemDescription, req.ItemQuantity, req.ItemPrice)
	if err != nil {
		log.Printf("%s: failed to create order: %v", op, err)

		return responses.BadRequest(err)
	}

	if err := h.usecase.CreateOrder(r.Context(), user); err != nil {
		log.Printf("%s: failed to create order: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.Created()

}
