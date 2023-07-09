package entities

import (
	"time"

	"github.com/google/uuid"
)

type Number struct {
	DBModel
	No               string    `gorm:"Column:number" json:"number"`
	UserId           uuid.UUID `gorm:"type:uuid" json:"user_id"` // shared number owned to lets_go
	IsShared         bool      `gorm:"Column:is_shared"`
	Subscribed       bool      `gorm:"Column:subscribed"`
	IsActive         bool      `gorm:"Column:is_active"`
	SubscriptionDate time.Time `json:"subscription_date"`
}
