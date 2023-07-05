package service

import (
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/google/uuid"
)

type AdminService struct {
	db ports.AdminActionsRepositoryContract
	pv ports.AdminActionsProviderContract
}

func NewAdminService() *AdminService {
	db := repositories.NewGormDatabase()
	// define pv
	return &AdminService{
		db: db,
		// set pv
	}
}

func (svc *AdminService) IsAdmin(userId uuid.UUID) error {
	// to implement with provider
	return nil
}

func (svc *AdminService) EditSingleMessagePrice(userId uuid.UUID, price int) error {
	if err := svc.IsAdmin(userId); err != nil {
		return err
	}

	err := svc.db.EditSingleMessagePrice(price)
	if err != nil {
		return err
	}

	return nil
}

func (svc *AdminService) EditGroupMessagePrice(userId uuid.UUID, price int) error {
	if err := svc.IsAdmin(userId); err != nil {
		return err
	}

	err := svc.db.EditGroupMessagePrice(price)
	if err != nil {
		return err
	}

	return nil
}

func (svc *AdminService) DisableUserAccount(userId uuid.UUID, target uuid.UUID) error {
	if err := svc.IsAdmin(userId); err != nil {
		return err
	}

	err := svc.pv.DisableUserAccount(target)
	if err != nil {
		return err
	}

	return nil
}

func (svc *AdminService) GetUserHistory(userId uuid.UUID, target uuid.UUID) ([]entities.Message, error) {
	if err := svc.IsAdmin(userId); err != nil {
		return []entities.Message{}, err
	}

	messages, err := svc.db.GetUserHistory(target)
	if err != nil {
		return []entities.Message{}, err
	}

	return messages, nil
}
