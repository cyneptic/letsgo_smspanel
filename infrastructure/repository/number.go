package repositories

import "github.com/google/uuid"

func (r *PGRepository) BuyANumber(userID, numberID uuid.UUID) error {
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
