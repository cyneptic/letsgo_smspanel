package entities

import "github.com/google/uuid"

type Template struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;" json:"id"`
	TempName string    `gorm:"tempname" json:"temp"`
	Content  string    `gorm:"content" json:"content"`
}
