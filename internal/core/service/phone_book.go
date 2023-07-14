package service

import (
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
)

type PhoneBookService struct {
	db ports.PhoneBookRepositoryContract
}

func NewPhoneBookService() *PhoneBookService {
	db := repositories.NewGormDatabase()

	return &PhoneBookService{
		db: db,
	}
}

func (svc *PhoneBookService) CreatePhoneBook(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {
	return svc.db.CreatePhoneBook(phoneBookModel)
}

func (svc *PhoneBookService) GetPhoneBookList(phoneBookModel entities.PhoneBook) ([]entities.PhoneBook, error) {
	return svc.db.GetPhoneBookList(phoneBookModel)
}

func (svc *PhoneBookService) GetPhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {
	return svc.db.GetPhoneBookById(phoneBookModel)
}

func (svc *PhoneBookService) UpdatePhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error) {
	return svc.db.UpdatePhoneBookById(phoneBookModel)
}

func (svc *PhoneBookService) DeletePhoneBookById(phoneBookModel entities.PhoneBook) error {
	return svc.db.DeletePhoneBookById(phoneBookModel)
}
