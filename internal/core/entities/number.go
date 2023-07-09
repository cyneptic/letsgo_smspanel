package entities

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type NumberType string

func (nt *NumberType) Scan(value interface{}) error {
	*nt = NumberType(value.([]byte))
	return nil
}

func (nt NumberType) Value() (driver.Value, error) {
	return string(nt), nil
}

const (
	SHARED     NumberType = "Shared"
	SUBSCRIBED            = "Subscribed"
	BOUGHT                = "Bought"
)

type Number struct {
	DBModel
	No   string     `gorm:"Column:number" json:"number"`
	Type NumberType `gorm:"type:number_type"`
}

type UserNumbers struct {
	DBModel
	UserId           uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Number           uuid.UUID `gorm:"type:uuid" json:"number"`
	SubscriptionDate time.Time `json:"subscription_date"`
}
