package entities

import "github.com/google/uuid"

type Message struct {
	DBModel
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Content  string    `json:"content"`
	Receiver string    `json:"receiver"`
	Sender   string    `json:"sender"`
}
