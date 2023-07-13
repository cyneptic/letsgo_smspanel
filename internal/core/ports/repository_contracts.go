package ports

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

type TemplateRepositoryContract interface {
	AddTemplate(temp entities.Template) error
	GetTemplate(tempname string) (entities.Template, error)
	AllTemplates() ([]entities.Template, error)
}

type SnedSMSRepositoryContract interface {
	RequestContactList(id uuid.UUID) ([]entities.Contact, error)
	RequestNumber(id uuid.UUID) (entities.Number, error)
	RequestUser(id uuid.UUID) (entities.User, error)
	WithdrawFromWallet(userid uuid.UUID, amount int) error
	GetSinglePrice() (int, error)
	GetGroupPrice() (int, error)
}
type PhoneBookRepositoryContract interface {
	CreatePhoneBookList(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	GetPhoneBookList(phoneBookModel entities.PhoneBook) ([]entities.PhoneBook, error)
	GetPhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	UpdatePhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	DeletePhoneBookById(phoneBookModel entities.PhoneBook) error
}

type ContactRepositoryContract interface {
	CreateContact(contactModel entities.Contact) (entities.Contact, error)
	GetContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	DeleteContactByUsername(contactModel entities.Contact) error
	GetContactList(contactModel entities.Contact) ([]entities.Contact, error)
	GetContactById(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactById(contactModel entities.Contact) (entities.Contact, error)
	DeleteContactById(contactModel entities.Contact) error
}

type NumberRepositoryContract interface {
	BuyANumber(userID uuid.UUID, number string) error
	GetSharedANumber() ([]entities.Number, error)
	IsReserved(number string) (bool, error)
	SubscribeMe(user uuid.UUID, number string) error
	WithdrawFromWallet(userid uuid.UUID, amount int) error
}

type AdminActionsRepositoryContract interface {
	EditSingleMessagePrice(amount int) error
	EditGroupMessagePrice(amount int) error
	GetUserHistory(uId uuid.UUID) ([]entities.Message, error)
	SearchAllMessages(query string) ([]entities.Message, error)
	AddBlacklistWord(word string) error
	RemoveBlacklistWord(word string) error
	AddBlacklistRegex(regex string) error
	RemoveBlacklistRegex(regex string) error
}

type PaymentDbContract interface {
	SetPaymentRequest(orderID uuid.UUID, payerID, refID, amount string) error
	VerifyPaymentRequest(payerID, refID, orderID string) (string, bool, error)
}
