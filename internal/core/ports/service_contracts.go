package ports

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"time"
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

type NumberServiceContract interface {
	GenerateNumber() (string, error)
	BuyNumber(user uuid.UUID, number string) error
	SubscribeNumber(user uuid.UUID, number string) error
	GetSharedNumber() ([]string, error)
}
type AdminActionServiceContract interface {
	EditSingleMessagePrice(userId uuid.UUID, price int) error
	EditGroupMessagePrice(userId uuid.UUID, price int) error
	DisableUserAccount(userId uuid.UUID, target uuid.UUID, toggle bool) error
	GetUserHistory(userId uuid.UUID, target uuid.UUID) ([]entities.Message, error)
	SearchAllMessages(userid uuid.UUID, query string) ([]entities.Message, error)
	AddBlacklistWord(userid uuid.UUID, word string) error
	RemoveBlacklistWord(userid uuid.UUID, word string) error
	AddBlacklistRegex(userid uuid.UUID, regex string) error
	RemoveBlacklistRegex(userid uuid.UUID, regex string) error
}
