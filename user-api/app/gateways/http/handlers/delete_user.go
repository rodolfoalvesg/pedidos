package handlers

import (
	"log"
	"net/http"
	"user-api/app/gateways/http/handlers/errors"
	"user-api/app/gateways/http/rest/requests"
	"user-api/app/gateways/http/rest/responses"
)

// DeleteUser deletes a user
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept json
//
// @Param user-id path string true "User ID"
//
// @Success 204 "No Content"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/user-api/users/{user-id} [delete]
func (h *Handler) DeleteUser(r *http.Request) responses.Response {
	const op = "UserHandler.DeleteUser"

	userID, err := requests.ParseUUID(r, "user-id")
	if err != nil {
		log.Printf("%s: failed to parse user ID: %v", op, err)

		return responses.BadRequest(err)
	}

	if err := h.usecase.DeleteUser(r.Context(), userID); err != nil {
		log.Printf("%s: failed to delete user: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.NoContent()
}
