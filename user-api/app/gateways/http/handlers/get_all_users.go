package handlers

import (
	"log"
	"net/http"
	"user-api/app/gateways/http/handlers/schema"
	"user-api/app/gateways/http/rest/requests"
	"user-api/app/gateways/http/rest/responses"
)

// GetAllUsers gets all users
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
//
// @Param public-id query string false "Public ID"
// @Param name query string false "Name"
//
// @Success 200 {object} []schema.UserResponse "Success"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/user-api/users [get]
func (h *Handler) GetAllUsers(r *http.Request) responses.Response {
	const op = "UserHandler.GetAllUsers"

	filter := requests.ParseQueryParams(r)

	users, err := h.usecase.GetAllUsers(r.Context(), filter)
	if err != nil {
		log.Printf("%s: failed to get all users: %v", op, err)

		return responses.InternalServerError(err)
	}

	return responses.OK(schema.MapUsersToResponse(users))
}
