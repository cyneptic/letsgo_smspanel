package repositories

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

func (db *PGRepository) EditSingleMessagePrice(amount int) error {
	if err := db.DB.Model(&entities.Prices{}).Update("single", amount).Error; err != nil {
		return err
	}

	return nil
}

func (db *PGRepository) EditGroupMessagePrice(amount int) error {
	if err := db.DB.Model(&entities.Prices{}).Update("group", amount).Error; err != nil {
		return err
	}

	return nil
}

func (db *PGRepository) GetUserHistory(uId uuid.UUID) ([]entities.Message, error) {
	var history []entities.Message
	if err := db.DB.Model(&entities.Message{}).Where("user_id = ?", uId).Find(&history).Error; err != nil {
		return []entities.Message{}, err
	}

	return history, nil
}
