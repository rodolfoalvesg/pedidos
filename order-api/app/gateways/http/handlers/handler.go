package handlers

import "order-api/app/domain/usecases"

type Handler struct {
	usecase usecases.OrderUsecase
}

func NewHandler(usecase usecases.OrderUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
