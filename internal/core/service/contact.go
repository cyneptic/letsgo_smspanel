package service

import (
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
)

type ContactService struct {
	db ports.ContactRepositoryContract
}

func NewContactService() *ContactService {
	db := repositories.NewGormDatabase()

	return &ContactService{
		db: db,
	}
}

func (svc *ContactService) CreateContact(contactModel entities.Contact) (entities.Contact, error) {
	return svc.db.CreateContact(contactModel)
}

func (svc *ContactService) GetContactByUsername(contactModel entities.Contact) (entities.Contact, error) {
	return svc.db.GetContactByUsername(contactModel)
}

func (svc *ContactService) UpdateContactByUsername(contactModel entities.Contact) (entities.Contact, error) {
	return svc.db.UpdateContactById(contactModel)
}

func (svc *ContactService) DeleteContactByUsername(contactModel entities.Contact) error {
	return svc.db.DeleteContactById(contactModel)
}

func (svc *ContactService) GetContactList(contactModel entities.Contact) ([]entities.Contact, error) {
	return svc.db.GetContactList(contactModel)
}

func (svc *ContactService) GetContactById(contactModel entities.Contact) (entities.Contact, error) {
	return svc.db.GetContactById(contactModel)
}

func (svc *ContactService) UpdateContactById(contactModel entities.Contact) (entities.Contact, error) {
	return svc.db.UpdateContactById(contactModel)
}

func (svc *ContactService) DeleteContactById(contactModel entities.Contact) error {
	return svc.db.DeleteContactById(contactModel)
}
