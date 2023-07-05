package entities

import (
	"time"
)

type User struct {
	DBModel
	Name        string    `gorm:"size:255" json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `gorm:"unique;size:255" json:"phone_number"`
	Email       string    `gorm:"unique;size:255" json:"email"`
	Password    string    `gorm:"size:255" json:"password"`
	Role        string    `gorm:"size:255" json:"role"`
}
