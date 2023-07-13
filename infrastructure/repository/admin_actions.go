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

func (db *PGRepository) GetSinglePrice() (int, error) {
	var prices entities.Prices
	if err := db.DB.Model(&entities.Prices{}).First(&prices).Error; err != nil {
		return 0, err
	}

	return prices.SingleMessage, nil
}

func (db *PGRepository) GetGroupPrice() (int, error) {
	var prices entities.Prices
	if err := db.DB.Model(&entities.Prices{}).First(&prices).Error; err != nil {
		return 0, err
	}

	return prices.GroupMessage, nil
}

func (db *PGRepository) GetUserHistory(uId uuid.UUID) ([]entities.Message, error) {
	var history []entities.Message
	if err := db.DB.Model(&entities.Message{}).Where("user_id = ?", uId).Find(&history).Error; err != nil {
		return []entities.Message{}, err
	}

	return history, nil
}

func (db *PGRepository) SearchAllMessages(query string) ([]entities.Message, error) {
	var result []entities.Message
	if err := db.DB.Model(&entities.Message{}).Where("content LIKE ?", "%"+query+"%").Find(&result).Error; err != nil {
		return []entities.Message{}, err
	}

	return result, nil
}

func (db *PGRepository) AddBlacklistWord(word string) error {
	q := entities.BlacklistWord{
		Word: word,
	}
	if err := db.DB.Create(&q).Error; err != nil {
		return err
	}

	return nil
}

func (db *PGRepository) RemoveBlacklistWord(word string) error {
	if err := db.DB.Delete(&entities.BlacklistWord{}, "word = ?", word).Error; err != nil {
		return err
	}

	return nil
}

func (db *PGRepository) AddBlacklistRegex(regex string) error {
	q := entities.BlacklistRegex{
		Expression: regex,
	}
	if err := db.DB.Create(&q).Error; err != nil {
		return err
	}

	return nil
}

func (db *PGRepository) RemoveBlacklistRegex(regex string) error {
	if err := db.DB.Delete(&entities.BlacklistRegex{}, "expression = ?", regex).Error; err != nil {
		return err
	}

	return nil
}
