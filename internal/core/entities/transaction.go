package entities

import "github.com/google/uuid"

type Transaction struct {
	DBModel
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	WalletID uuid.UUID `gorm:"type:uuid" json:"wallet_id"`
	Amount   int       `json:"amount"`
	Status   string    `json:"status"`
}
