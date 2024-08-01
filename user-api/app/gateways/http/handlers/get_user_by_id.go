package handlers

import (
	"log"
	"net/http"
	"user-api/app/gateways/http/handlers/errors"
	"user-api/app/gateways/http/handlers/schema"
	"user-api/app/gateways/http/rest/requests"
	"user-api/app/gateways/http/rest/responses"
)

// GetUserByID gets a user by ID
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept json
//
// @Param user-id path string true "User ID"
//
// @Success 200 {object} schema.UserResponse "Success"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 404 {object} responses.NotFoundError "Not Found"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/user-api/users/{user-id} [get]
func (h *Handler) GetUserByID(r *http.Request) responses.Response {
	const op = "UserHandler.GetUserByID"

	userID, err := requests.ParseUUID(r, "user-id")
	if err != nil {
		log.Printf("%s: failed to parse user ID: %v", op, err)

		return responses.BadRequest(err)
	}

	user, err := h.usecase.GetUserByID(r.Context(), userID)
	if err != nil {
		log.Printf("%s: failed to get user by email: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.OK(schema.MapUserToResponse(*user))
}
