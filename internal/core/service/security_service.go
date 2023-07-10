package service

import (
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
)

type SecurityService struct {
	db ports.SecurityRepositoryContract
}

func NewSecurityService() *SecurityService {
	db := repositories.NewGormDatabase()
	return &SecurityService{
		db: db,
	}
}

// search string in messages
func (r *SecurityService) SearchInMessages(word string) ([]entities.Message, error) {
	messages, err := r.db.SearchInMessages(word)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// done
func (r *SecurityService) AddWordToBlackList(word entities.BlacklistWord) error {

	err := r.db.AddWordToBlackList(word)

	if err != nil {
		return err
	}

	return nil
}

// done
func (r *SecurityService) RemoveWordFromBlackList(word string) error {

	err := r.db.RemoveWordFromBlackList(word)
	if err != nil {
		return err
	}
	return nil
}

// done
func (r *SecurityService) AddRegexToBlackList(regex entities.BlacklistRegex) error {
	err := r.db.AddRegexToBlackList(regex)
	if err != nil {
		return err
	}
	return nil
}

// done
func (r *SecurityService) RemoveRegexFromBlackList(regex string) error {
	err := r.db.RemoveRegexFromBlackList(regex)
	if err != nil {
		return err
	}
	return nil
}
