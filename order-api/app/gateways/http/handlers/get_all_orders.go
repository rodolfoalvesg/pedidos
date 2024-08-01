package handlers

import (
	"log"
	"net/http"
	"order-api/app/gateways/http/handlers/schema"
	"order-api/app/gateways/http/rest/requests"
	"order-api/app/gateways/http/rest/responses"
)

// GetAllOrders gets all ordes
// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Accept json
//
// @Param order-id query string false "Order ID"
// @Param name query string false "Name"
// @Param user-id query string false "User ID"
//
// @Success 200 {object} []schema.OrderResponse "Success"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/order-api/orders [get]
func (h *Handler) GetAllOrders(r *http.Request) responses.Response {
	const op = "OrderHandler.GetAllOrders"

	filter := requests.ParseQueryParams(r)

	orders, err := h.usecase.GetAllOrders(r.Context(), filter)
	if err != nil {
		log.Printf("%s: failed to get all orders: %v", op, err)

		return responses.InternalServerError(err)
	}

	return responses.OK(schema.MapOrdersToResponse(orders))
}
