package errors

import (
	"errors"
	"user-api/app/domain/entities"
	"user-api/app/gateways/http/rest/responses"
)

var (
	ErrInvalidRequestID = errors.New("invalid request id")
)

// SelectError select response custom error
func SelectError(err error) responses.Response {
	switch {
	case errors.Is(err, entities.ErrorUserNotFound):
		return responses.NotFound(entities.ErrorUserNotFound)
	case errors.Is(err, entities.ErrUserAlreadyExists):
		return responses.Conflict(entities.ErrUserAlreadyExists)
	default:
		return responses.InternalServerError(err)
	}
}
