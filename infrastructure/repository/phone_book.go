package repositories

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func (r *PGRepository) CreatePhoneBook(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {

	var phoneBook entities.PhoneBook
	phoneBook.Name = phoneBookModel.Name
	phoneBook.UserId = phoneBookModel.UserId

	// Create phone book entry
	err := r.DB.Create(&phoneBook).Error
	if err != nil {
		return phoneBook, err
	}

	return phoneBook, nil
}

func (r *PGRepository) GetPhoneBookList(phoneBookModel entities.PhoneBook) ([]entities.PhoneBook, error) {

	phoneBooks := []entities.PhoneBook{}

	// Retrieve phone book list
	err := r.DB.Where("user_id = ?", phoneBookModel.UserId).Find(&phoneBooks).Error
	if err != nil {
		return phoneBooks, err
	}

	return phoneBooks, nil
}

func (r *PGRepository) GetPhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {

	var phoneBook entities.PhoneBook

	// Retrieve phone book by ID
	err := r.DB.Where("user_id = ? AND id = ?", phoneBookModel.UserId, phoneBookModel.ID).First(&phoneBook).Error
	if err != nil {
		return entities.PhoneBook{}, err
	}

	return phoneBook, nil
}

func (r *PGRepository) UpdatePhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {

	var phoneBook entities.PhoneBook

	// Retrieve phone book by ID
	err := r.DB.Where("user_id = ? AND id = ?", phoneBookModel.UserId, phoneBookModel.ID).First(&phoneBook).Error
	if err != nil {
		return phoneBook, err
	}

	phoneBook.Name = phoneBookModel.Name
	phoneBook.ModifiedAt = time.Now()

	err = r.DB.Save(&phoneBook).Error
	if err != nil {
		return phoneBook, err
	}

	return phoneBook, nil
}

func (r *PGRepository) DeletePhoneBookById(phoneBookModel entities.PhoneBook) error {

	var phoneBook entities.PhoneBook

	// Delete phone book by ID
	err := r.DB.Where("user_id = ? AND id = ?", phoneBookModel.UserId, phoneBookModel.ID).Delete(&phoneBook).Error
	if err != nil {
		return err
	}

	return nil
}
