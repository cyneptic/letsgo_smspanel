package entities

import (
	"github.com/google/uuid"
)

type PhoneBook struct {
	DBModel
	UserId uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Name   string    `gorm:"size:255" json:"name"`
}
