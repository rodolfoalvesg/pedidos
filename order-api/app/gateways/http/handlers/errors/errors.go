package errors

import (
	"errors"
	"order-api/app/domain/entities"
	"order-api/app/gateways/http/rest/responses"
)

var (
	ErrInvalidRequestID = errors.New("invalid request id")
)

// SelectError select response custom error
func SelectError(err error) responses.Response {
	switch {
	case errors.Is(err, entities.ErrorOrderNotFound):
		return responses.NotFound(entities.ErrorOrderNotFound)
	default:
		return responses.InternalServerError(err)
	}
}
