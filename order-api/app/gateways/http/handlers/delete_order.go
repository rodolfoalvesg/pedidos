package handlers

import (
	"log"
	"net/http"
	"order-api/app/gateways/http/handlers/errors"
	"order-api/app/gateways/http/rest/requests"
	"order-api/app/gateways/http/rest/responses"
)

// DeleteOrder deletes an order.
// @Summary Delete a order
// @Description Delete a order
// @Tags orders
// @Accept json
//
// @Param order-id path string true "Order ID"
//
// @Success 204 "No Content"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/order-api/orders/{order-id} [delete]
func (h *Handler) DeleteOrder(r *http.Request) responses.Response {
	const op = "OrderHandler.DeleteOrder"

	orderID, err := requests.ParseUUID(r, "order-id")
	if err != nil {
		log.Printf("%s: failed to parse order ID: %v", op, err)

		return responses.BadRequest(err)
	}

	if err := h.usecase.DeleteOrder(r.Context(), orderID); err != nil {
		log.Printf("%s: failed to delete order: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.NoContent()
}
