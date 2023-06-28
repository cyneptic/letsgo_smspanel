package entities

import "github.com/google/uuid"

type Wallet struct {
	DBModel
	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Credit int
}
