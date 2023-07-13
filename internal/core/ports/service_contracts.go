package ports

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

type TemplateContract interface {
	CreateTemplate(temp entities.Template) error
	GetTemplateMapContent(tempName string) (string, map[string]string, error)
	GenerateTemplate(content string, temp map[string]string) (string, error)
	GetAllTemplates() ([]entities.Template, error)
}

type SendSMSServiceContract interface {
	SendToContactList(msg entities.Message) error
	SendToNumber(msg entities.Message) error
	SendToUser(msg entities.Message) error
	SendToContactListInterval(msg entities.Message, internal time.Duration) error
}

type PhoneBookServiceContract interface {
	CreatePhoneBook(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	GetPhoneBookList(phoneBookModel entities.PhoneBook) ([]entities.PhoneBook, error)
	GetPhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	UpdatePhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	DeletePhoneBookById(phoneBookModel entities.PhoneBook) error
}

type ContactServiceContract interface {
	CreateContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	ListContactByUsername(contactModel entities.Contact) ([]entities.Contact, error)
	GetContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactByUsername(username string,contactModel entities.Contact) (entities.Contact, error)
	DeleteContactByUsername(contactModel entities.Contact) error
	
	CreateContact(contactModel entities.Contact) (entities.Contact, error)
	GetContactList(contactModel entities.Contact) ([]entities.Contact, error)
	GetContactById(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactById(contactModel entities.Contact) (entities.Contact, error)
	DeleteContactById(contactModel entities.Contact) error
}
