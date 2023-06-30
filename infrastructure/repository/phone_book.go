package repositories

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func (r *PGRepository) CreateUser(phoneBookModel entities.User) (entities.User, error) {
	var User entities.User
	User.Name = phoneBookModel.Name
	User.Email = phoneBookModel.Email
	User.Password = phoneBookModel.Password
	User.PhoneNumber = phoneBookModel.PhoneNumber
	User.Role = phoneBookModel.Role

	err := r.DB.Create(&User).Error
	return User, err
}
func (r *PGRepository) CreatePhoneBookList(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {
	var PhoneBook entities.PhoneBook
	PhoneBook.Name = phoneBookModel.Name
	PhoneBook.UserId = phoneBookModel.UserId

	err := r.DB.Create(&PhoneBook).Error
	return PhoneBook, err
}

func (r *PGRepository) GetPhoneBookList(phoneBookModel entities.PhoneBook) ([]entities.PhoneBook, error) {
	PhoneBooks := []entities.PhoneBook{}

	err := r.DB.Where("user_id = ?", phoneBookModel.UserId).Find(&PhoneBooks).Error
	if err != nil {
		return PhoneBooks, err
	}

	return PhoneBooks, nil
}

func (r *PGRepository) GetPhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {
	var phoneBook entities.PhoneBook

	err := r.DB.Where("user_id = ? AND id = ?", phoneBookModel.UserId, phoneBookModel.ID).First(&phoneBook).Error
	if err != nil {
		return entities.PhoneBook{}, err
	}

	return phoneBook, nil
}

func (r *PGRepository) UpdatePhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {
	var phoneBook entities.PhoneBook

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

	err := r.DB.Where("user_id = ? AND id = ?", phoneBookModel.UserId, phoneBookModel.ID).Delete(&phoneBook).Error
	if err != nil {
		return err
	}

	return nil
}
