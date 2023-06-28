package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	Role        string    `gorm:"type:ENUM('user','admin','superAdmin');default:'user'" json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time
	DeletedAt   time.Time
}
