package repositories

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
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
func (r *PGRepository) IsSubscribable(user uuid.UUID, number string) (bool, error) {
	return true, nil
}
func (r *PGRepository) SubscribeMe(user uuid.UUID, number string) error {
	var Numb entities.Number
	Numb.ID = uuid.New()
	Numb.No = number
	Numb.UserId = user
	Numb.IsShared = false
	Numb.Subscribed = true
	Numb.IsActive = true
	err := r.DB.Create(&Numb).Error
	return err
}

func (r *PGRepository) IsReserved(randomNumber string) bool {
	return false
}
