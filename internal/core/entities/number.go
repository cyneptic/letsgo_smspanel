package entities

import (
	"time"

	"github.com/google/uuid"
)

type Number struct {
	DBModel
	No               string    `gorm:"Column:number" json:"number"`
	UserId           uuid.UUID `gorm:"Column:user_id;type:uuid" json:"user_id"`
	Shared           bool      `json:"shared"`
	Subscription     bool      `json:"subscription"`
	SubscriptionDate time.Time `json:"subscription_date"`
}
