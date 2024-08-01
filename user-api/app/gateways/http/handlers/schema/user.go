package schema

import (
	"time"
	"user-api/app/domain/entities"

	"github.com/google/uuid"
)

type UserRequest struct {
	Name  string `json:"name"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserResponse struct {
	PublicID  uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapUserToResponse(user entities.User) UserResponse {
	return UserResponse{
		PublicID:  user.PublicID,
		Name:      user.Name,
		CPF:       user.CPF,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func MapUsersToResponse(users []*entities.User) []entities.User {
	var usersResponse []entities.User
	for _, user := range users {
		usersResponse = append(usersResponse, *user)
	}

	return usersResponse
}
