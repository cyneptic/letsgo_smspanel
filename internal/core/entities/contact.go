package entities

import "github.com/google/uuid"

type Contact struct {
	DBModel
	PhoneBookID uuid.UUID `json:"phone_book_id"`
	FirstName   string    `gorm:"size:255" json:"first_name"`
	LastName    string    `gorm:"size:255" json:"last_name"`
	Username    string    `gorm:"size:255;" json:"user_name"` // Has to be unique based on phonebook
	PhoneNumber string    `json:"phone_number"`
}
