package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	DBModel
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	Role        string    `gorm:"default:'user'" json:"role"`
}
