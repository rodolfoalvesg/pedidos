package responses

import "net/http"

type Response struct {
	Status  int
	Error   error
	header  http.Header
	Payload interface{}
}

func (r Response) Header() http.Header {
	return r.header
}

func OK(payload interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Payload: payload,
	}
}

func Created() Response {
	return Response{
		Status: http.StatusCreated,
	}
}

func BadRequest(err error) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Error:   err,
		Payload: ErrBadRequest,
	}
}

func InternalServerError(err error) Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Error:   err,
		Payload: ErrInternalServerError,
	}
}

func NotFound(err error) Response {
	return Response{
		Status:  http.StatusNotFound,
		Error:   err,
		Payload: ErrNotFound,
	}
}

func Conflict(err error) Response {
	return Response{
		Status:  http.StatusConflict,
		Error:   err,
		Payload: ErrConflict,
	}
}

func NoContent() Response {
	return Response{
		Status: http.StatusNoContent,
	}
}
