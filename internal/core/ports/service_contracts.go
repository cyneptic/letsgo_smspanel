package ports

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

type PhoneBookServiceContract interface {
	CreatePhoneBookList(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	GetPhoneBookList(phoneBookModel entities.PhoneBook) ([]entities.PhoneBook, error)
	GetPhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	UpdatePhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	DeletePhoneBookById(phoneBookModel entities.PhoneBook) error
}

type ContactServiceContract interface {
	CreateContact(contactModel entities.Contact) (entities.Contact, error)
	GetContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	DeleteContactByUsername(contactModel entities.Contact) error
	GetContactList(contactModel entities.Contact) ([]entities.Contact, error)
	GetContactById(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactById(contactModel entities.Contact) (entities.Contact, error)
	DeleteContactById(contactModel entities.Contact) error
}

type AdminActionServiceContract interface {
	EditSingleMessagePrice(userId uuid.UUID, price int) error
	EditGroupMessagePrice(userId uuid.UUID, price int) error
	DisableUserAccount(userId uuid.UUID, target uuid.UUID, toggle bool) error
	GetUserHistory(userId uuid.UUID, target uuid.UUID) ([]entities.Message, error)
}
