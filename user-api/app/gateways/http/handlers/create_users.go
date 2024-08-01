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

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
//
// @Param Body body schema.UserRequest true "User Request"
//
// @Success 201 "Created"
// @Failure 400 {object} responses.BadRequestError "Bad Request"
// @Failure 500 {object} responses.InternalServerErr "Internal Server Error"
// @Router /api/v1/user-api/users [post]
func (h *Handler) CreateUser(r *http.Request) responses.Response {
	const op = "UserHandler.CreateUser"

	var req schema.UserRequest
	if err := requests.DecodeBodyJSON(r, &req); err != nil {
		log.Printf("%s: failed to decode body JSON: %v", op, err)

		return responses.BadRequest(err)
	}

	user, err := entities.NewUser(req.Name, req.CPF, req.Email, req.Phone)
	if err != nil {
		log.Printf("%s: failed to create user: %v", op, err)

		return responses.BadRequest(err)
	}

	if err := h.usecase.CreateUser(r.Context(), user); err != nil {
		log.Printf("%s: failed to create user: %v", op, err)

		return errors.SelectError(err)
	}

	return responses.Created()

}
