package repositories

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"time"
)

func (r *PGRepository) BuyANumber(userID uuid.UUID, number string) error {
	var Num entities.UserNumbers
	var no entities.Number
	no.ID = uuid.New()
	no.No = number
	no.Type = entities.BOUGHT
	Num.ID = uuid.New()
	Num.UserId = userID
	Num.Number = no.ID
	Num.SubscriptionDate = time.Now()
	r.DB.Create(&Num)
	return nil
}

func (r *PGRepository) GetShareANumber() (string, error) {
	return "", nil
}

func (r *PGRepository) IsNumberFree(number string) (bool, error) {
	return true, nil
}
func (r *PGRepository) IsSubscribable(user, number string) (bool, error) {
	return true, nil
}
func (r *PGRepository) SubscribeMe(user, number string) {

}

func (r *PGRepository) IsReserved(randomNumber string) bool {
	return false
}
