package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired        = errors.New("name is required")
	ErrCPFIsRequired         = errors.New("cpf is required")
	ErrEmailIsRequired       = errors.New("email is required")
	ErrorUserNotFound        = errors.New("user not found")
	ErrorUserNotFoundInCache = errors.New("user not found in cache")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrCPFAlreadyExists      = errors.New("cpf already exists")
	ErrUserAlreadyExists     = errors.New("user already exists")
)

type User struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoIncrement:true"`
	PublicID  uuid.UUID `json:"public_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"not null"`
	CPF       string    `json:"cpf" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// NewUser creates a new User.
func NewUser(name, cpf, email, phone string) (*User, error) {
	user := User{
		Name:  name,
		CPF:   cpf,
		Email: email,
		Phone: phone,
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return &user, nil
}

// validate validates the User entities.
func (u *User) validate() error {
	if u.Name == "" {
		return ErrNameIsRequired
	}

	if u.CPF == "" {
		return ErrCPFIsRequired
	}

	if u.Email == "" {
		return ErrEmailIsRequired
	}

	return nil
}
