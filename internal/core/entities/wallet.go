package entities

import "github.com/google/uuid"

type Wallet struct {
	DBModel
	UserID uuid.UUID `gorm:"type:uuid;unique" json:"user_id"`
	Credit int `json:"credit"`
}
