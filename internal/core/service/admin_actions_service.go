package service

import (
	"errors"
	"regexp"
	"strings"

	"github.com/cyneptic/letsgo-smspanel/infrastructure/client"
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
	pv := client.NewAdminActionClient()
	return &AdminService{
		db: db,
		pv: pv,
	}
}
func (svc *AdminService) SearchAllMessages(userid uuid.UUID, query string) ([]entities.Message, error) {
	if err := svc.IsAdmin(userid); err != nil {
		return []entities.Message{}, err
	}

	messages, err := svc.db.SearchAllMessages(query)
	if err != nil {
		return []entities.Message{}, err
	}

	result := svc.FilterSensitiveInformation(messages)

	return result, nil
}

func (svc *AdminService) AddBlacklistWord(userid uuid.UUID, word string) error {
	if err := svc.IsAdmin(userid); err != nil {
		return err
	}
	if word == "" {
		return errors.New("cannot blacklist empty")
	}

	err := svc.db.AddBlacklistWord(word)
	if err != nil {
		return err
	}

	return nil
}

func (svc *AdminService) RemoveBlacklistWord(userid uuid.UUID, word string) error {
	if err := svc.IsAdmin(userid); err != nil {
		return err
	}

	if word == "" {
		return errors.New("cannot blacklist empty")
	}

	err := svc.db.RemoveBlacklistWord(word)
	if err != nil {
		return err
	}

	return nil

}

func (svc *AdminService) AddBlacklistRegex(userid uuid.UUID, regex string) error {
	if err := svc.IsAdmin(userid); err != nil {
		return err
	}

	err := svc.db.AddBlacklistRegex(regex)
	if err != nil {
		return err
	}

	return nil

}

func (svc *AdminService) RemoveBlacklistRegex(userid uuid.UUID, regex string) error {
	if err := svc.IsAdmin(userid); err != nil {
		return err
	}

	err := svc.db.RemoveBlacklistRegex(regex)
	if err != nil {
		return err
	}

	return nil
}

var sensitiveKeywords = []string{
	"pincode",
	"password",
}

func (svc *AdminService) FilterSensitiveInformation(messages []entities.Message) []entities.Message {
	for _, message := range messages {
		for _, word := range sensitiveKeywords {
			if strings.Contains(message.Content, word) {
				exrp, _ := regexp.Compile(`\b\d{7,9}\b`)
				exrp.ReplaceAllString(message.Content, "_________")
			}
		}
	}

	return messages
}

func (svc *AdminService) IsAdmin(userId uuid.UUID) error {
	bool, err := svc.pv.IsAdmin(userId)
	if err != nil {
		return err
	}

	if !bool {
		return errors.New("You are not allowed to do this!")
	}

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

func (svc *AdminService) DisableUserAccount(userId uuid.UUID, target uuid.UUID, toggle bool) error {
	if err := svc.IsAdmin(userId); err != nil {
		return err
	}

	err := svc.pv.DisableUserAccount(target, toggle)
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

	result := svc.FilterSensitiveInformation(messages)

	return result, nil
}
