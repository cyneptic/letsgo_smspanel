package repositories

import (
	

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)


func (r *PGRepository) SearchInMessages(word string) ([]entities.Message, error) {
	var messages []entities.Message
	if err := r.DB.Where("content LIKE ?", "%"+word+"%").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}


func (r *PGRepository) AddWordToBlackList(word entities.BlacklistWord) error {
	if err := r.DB.Create(&word).Error; err != nil {
		return err
	}
	return nil
}


func (r *PGRepository) RemoveWordFromBlackList(word string) error {
	if err := r.DB.Where("word = ?", word).Unscoped().Delete(&entities.BlacklistWord{}).Error; err != nil {
		return err
	}
	return nil
}


func (r *PGRepository) AddRegexToBlackList(regex entities.BlacklistRegex) error {
	if err := r.DB.Create(&regex).Error; err != nil {
		return err
	}
	return nil
}

							// *** BUG *** //
func (r *PGRepository) RemoveRegexFromBlackList(regex string) error {
	
	if err := r.DB.Where("regex = ?", regex).Unscoped().Delete(&entities.BlacklistRegex{}).Error; err != nil {
		return err
	}
	return nil
}
