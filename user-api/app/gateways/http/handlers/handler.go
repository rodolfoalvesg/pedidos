package handlers

import "user-api/app/domain/usecases"

type Handler struct {
	usecase usecases.UserUsecase
}

func NewHandler(usecase usecases.UserUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
