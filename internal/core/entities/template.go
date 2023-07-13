package entities

import "github.com/google/uuid"

type Template struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;" json:"id"`
	TempName string    `gorm:"Column:tempname" json:"temp"`
	Content  string    `gorm:"Column:content" json:"content"`
}
