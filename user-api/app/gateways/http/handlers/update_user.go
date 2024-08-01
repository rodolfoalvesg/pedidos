package handlers

import (
	"log"
	"net/http"
	"user-api/app/domain/entities"
	"user-api/app/gateways/http/handlers/errors"
	"user-api/app/gateways/http/handlers/schema"
	"user-api/app/gateways/http/rest/requests"
	"user-api/app/gateways/http/rest/responses"
)

// UpdateUser creates a new user
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
//
// @Param user-id path string true "User ID"
// @Param Body body schema.UpdateUserRequest true "User Request Update"
//
// @Success 204 "No Content"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/user-api/users/{user-id} [put]
func (h *Handler) UpdateUser(r *http.Request) responses.Response {
	const op = "UserHandler.UpdateUser"

	userID, err := requests.ParseUUID(r, "user-id")
	if err != nil {
		log.Printf("%s: failed to parse userID: %v", op, err)

		return responses.BadRequest(err)
	}

	var req schema.UpdateUserRequest
	if err := requests.DecodeBodyJSON(r, &req); err != nil {
		log.Printf("%s: failed to decode body JSON: %v", op, err)

		return responses.BadRequest(err)
	}

	user := &entities.User{
		PublicID: userID,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	if err := h.usecase.UpdateUser(r.Context(), user); err != nil {
		log.Printf("%s: failed to update user: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.NoContent()
}
