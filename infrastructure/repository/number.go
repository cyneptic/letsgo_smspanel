package repositories

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"time"
)

func (r *PGRepository) BuyANumber(userID uuid.UUID, number string) error {
	var Numb entities.Number
	Numb.ID = uuid.New()
	Numb.No = number
	Numb.UserId = userID
	Numb.IsShared = false
	Numb.Subscribed = false
	Numb.IsActive = true
	err := r.DB.Create(&Numb).Error
	return err
}

func (r *PGRepository) GetSharedANumber() ([]entities.Number, error) {
	var sharedNumbers []entities.Number
	err := r.DB.Where("is_shared = ? AND is_active = ?", true, true).Find(&sharedNumbers).Error
	return sharedNumbers, err
}

func (r *PGRepository) IsNumberFree(number string) (bool, error) {
	return true, nil
}

func (r *PGRepository) SubscribeMe(user uuid.UUID, number string) error {
	var Numb entities.Number
	Numb.ID = uuid.New()
	Numb.No = number
	Numb.UserId = user
	Numb.IsShared = false
	Numb.Subscribed = true
	Numb.SubscriptionDate = time.Now().AddDate(0, 1, 0)
	Numb.IsActive = true
	err := r.DB.Create(&Numb).Error
	return err
}

func (r *PGRepository) IsReserved(number string) (bool, error) {
	var exists bool
	err := r.DB.Model(&entities.Number{}).Where("number = ?", number).Find(&exists).Error
	if exists || err != nil {
		return true, err
	}
	return false, nil
}
