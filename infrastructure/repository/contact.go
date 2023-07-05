package repositories

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func (r *PGRepository) CreateContact(contactModel entities.Contact) (entities.Contact, error) {
	contact := entities.Contact{}

	contact.UserID = contactModel.UserID
	contact.PhoneBookID = contactModel.PhoneBookID
	contact.FirstName = contactModel.FirstName
	contact.LastName = contactModel.LastName
	contact.PhoneNumber = contactModel.PhoneNumber
	contact.Username = contactModel.Username

	err := r.DB.Create(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (r *PGRepository) GetContactByUsername(contactModel entities.Contact) (entities.Contact, error) {
	var contact entities.Contact

	if err := r.DB.Where("user_id = ? AND Username = ?", contactModel.UserID, contactModel.Username).First(&contact).Error; err != nil {
		return contact, err
	}

	return contact, nil
}

func (r *PGRepository) UpdateContactByUsername(contactModel entities.Contact) (entities.Contact, error) {
	var contact entities.Contact

	err := r.DB.Where("user_id = ? AND username = ?", contactModel.UserID, contactModel.Username).First(&contact).Error
	if err != nil {
		return contact, err
	}

	if contactModel.FirstName != "" {
		contact.FirstName = contactModel.FirstName
	}

	if contactModel.LastName != "" {
		contact.LastName = contactModel.LastName
	}

	if contactModel.Username != "" {
		contact.Username = contactModel.Username
	}

	if contactModel.PhoneNumber != "" {
		contact.PhoneNumber = contactModel.PhoneNumber
	}
	contact.ModifiedAt = time.Now()

	err = r.DB.Save(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (r *PGRepository) DeleteContactByUsername(contactModel entities.Contact) error {
	var contact entities.Contact

	err := r.DB.Where("user_id = ? AND username = ?", contactModel.UserID, contactModel.Username).Delete(&contact).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PGRepository) GetContactList(contactModel entities.Contact) ([]entities.Contact, error) {
	contacts := []entities.Contact{}

	err := r.DB.Where("user_id = ? AND phone_book_id = ?", contactModel.UserID, contactModel.PhoneBookID).Find(&contacts).Error
	if err != nil {
		return contacts, err
	}

	return contacts, nil
}

func (r *PGRepository) GetContactById(contactModel entities.Contact) (entities.Contact, error) {
	var contact entities.Contact

	if err := r.DB.Where("user_id = ? AND id = ? AND phone_book_id = ?", contactModel.UserID, contactModel.ID, contactModel.PhoneBookID).First(&contact).Error; err != nil {
		return contact, err
	}

	return contact, nil
}

func (r *PGRepository) UpdateContactById(contactModel entities.Contact) (entities.Contact, error) {
	var contact entities.Contact

	err := r.DB.Where("user_id = ? AND id = ? AND phone_book_id = ?", contactModel.UserID, contactModel.ID, contactModel.PhoneBookID).First(&contact).Error
	if err != nil {
		return contact, err
	}

	if contactModel.FirstName != "" {
		contact.FirstName = contactModel.FirstName
	}

	if contactModel.LastName != "" {
		contact.LastName = contactModel.LastName
	}

	if contactModel.Username != "" {
		contact.Username = contactModel.Username
	}

	if contactModel.PhoneNumber != "" {
		contact.PhoneNumber = contactModel.PhoneNumber
	}
	contact.ModifiedAt = time.Now()

	err = r.DB.Save(&contact).Error
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (r *PGRepository) DeleteContactById(contactModel entities.Contact) error {
	var contact entities.Contact

	err := r.DB.Where("user_id = ? AND id = ? AND phone_book_id = ?", contactModel.UserID, contactModel.ID, contactModel.PhoneBookID).Delete(&contact).Error
	if err != nil {
		return err
	}

	return nil
}
