package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserIDIsRequired           = errors.New("user id is required")
	ErrOrderDescriptionIsRequired = errors.New("order description is required")
	ErrorItemQuantityIsInvalid    = errors.New("item quantity is invalid")
	ErrorItemPriceIsInvalid       = errors.New("item price is invalid")
	ErrorOrderNotFoundInCache     = errors.New("user not found in cache")
	ErrorOrderNotFound            = errors.New("order not found")
)

type Order struct {
	ID              uint      `json:"-" gorm:"primaryKey;autoIncrement:true"`
	PublicID        uuid.UUID `json:"public_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID          uuid.UUID `json:"user_id" gorm:"not null"`
	ItemDescription string    `json:"item_description" gorm:"not null"`
	ItemQuantity    int       `json:"item_quantity" gorm:"not null"`
	ItemPrice       float64   `json:"item_price" gorm:"not null"`
	TotalValue      float64   `json:"total_value" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// NewUser creates a new User.
func NewOrder(userID uuid.UUID, itemDescription string, itemQtd int, itemPrice float64) (*Order, error) {

	user := Order{
		UserID:          userID,
		ItemDescription: itemDescription,
		ItemQuantity:    itemQtd,
		ItemPrice:       itemPrice,
	}

	totalValue := (float64(itemQtd) * itemPrice)
	user.TotalValue = totalValue

	if err := user.validate(); err != nil {
		return nil, err
	}

	return &user, nil
}

// validate validates the User entities.
func (u *Order) validate() error {
	if u.UserID == uuid.Nil {
		return ErrUserIDIsRequired
	}

	if u.ItemDescription == "" {
		return ErrOrderDescriptionIsRequired
	}

	if u.ItemQuantity <= 0 {
		return ErrorItemQuantityIsInvalid
	}

	if u.ItemPrice <= 0 {
		return ErrorItemPriceIsInvalid
	}

	return nil
}
