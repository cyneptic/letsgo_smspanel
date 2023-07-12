package repositories

import (
	"errors"
	"fmt"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

// !send sms repository
func (pc *PGRepository) RequestContactList(id uuid.UUID) ([]entities.Contact, error) {
	var phoneBook entities.PhoneBook
	var contactlist []entities.Contact
	err := pc.DB.Model(&entities.PhoneBook{}).Where("user_id=?", id).First(&phoneBook).Error
	if err != nil {
		return []entities.Contact{}, errors.New("can't get phoneBook data from dataBase")
	}
	//! get phonebook
	fmt.Println("phonebookid", phoneBook.ID)
	err = pc.DB.Model(&entities.Contact{}).Where("phone_book_id=?", phoneBook.ID).Find(&contactlist).Error
	if err != nil {
		return []entities.Contact{}, errors.New("can't get contact data from dataBase")
	}
	return contactlist, nil

}

// !RequestNumber
func (pc *PGRepository) RequestNumber(id uuid.UUID) (entities.Number, error) {
	var number entities.Number
	err := pc.DB.Model(&entities.Number{}).Where("user_id=?", id).First(&number)
	if err != nil {
		return entities.Number{}, errors.New("can't get number data from database")
	}
	return number, nil
}

// !RequestUser
func (pc *PGRepository) RequestUser(id uuid.UUID) (entities.User, error) {
	var user entities.User
	err := pc.DB.Model(&entities.User{}).Where("id=?", id).First(&user).Error
	if err != nil {
		return entities.User{}, errors.New("can't get user data from database")
	}
	return user, nil
}
