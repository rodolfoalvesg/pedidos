package entities

type OrderFilter struct {
	PublicID    string `json:"public_id"`
	UserID      string `json:"user_id"`
	Description string `json:"description"`
}
