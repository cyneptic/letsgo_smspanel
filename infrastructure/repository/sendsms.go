package repositories

import (
	"errors"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

// !send sms repository
func (pc *PGRepository) RequestContactList() ([]entities.Contact, error) {
	var contactlist []entities.Contact
	err := pc.DB.Model(&entities.Contact{}).Find(&contactlist)
	if err != nil {
		return []entities.Contact{}, errors.New("can't get contact data from database")
	}
	return contactlist, nil

}

// !RequestNumber
func (pc *PGRepository) RequestNumber(id uuid.UUID) (entities.Number, error) {
	var number entities.Number
	err := pc.DB.Model(&entities.Number{}).Where("userid=?", id).First(&number)
	if err != nil {
		return entities.Number{}, errors.New("can't get number data from database")
	}
	return number, nil
}

// !RequestUser
func (pc *PGRepository) RequestUser(id uuid.UUID) (entities.User, error) {
	var user entities.User
	err := pc.DB.Model(&entities.User{}).Where("id=?", id).First(&user)
	if err != nil {
		return entities.User{}, errors.New("can't get user data from database")
	}
	return user, nil
}
