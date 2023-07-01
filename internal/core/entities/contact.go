package entities

import (
	"github.com/google/uuid"
)

type Contact struct {
	DBModel
	UserID      uuid.UUID `json:"user_id"`
	PhoneBookID uuid.UUID `json:"phone_book_id"`
	FirstName   string    `gorm:"size:255" json:"first_name"`
	LastName    string    `gorm:"size:255" json:"last_name"`
	Username    string    `gorm:"size:255;unique" json:"username"`
	PhoneNumber string    `json:"phone_number"`
}
