package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBModel struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;Column:id" json:"id"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  *gorm.DeletedAt
}

func (d *DBModel) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New()
	return nil
}
