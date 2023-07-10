package ports

import (
	

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

type SecurityRepositoryContract interface {
	SearchInMessages(word string) ([]entities.Message, error)
	AddWordToBlackList(word entities.BlacklistWord) error
	RemoveWordFromBlackList(word string) error
	AddRegexToBlackList(regex entities.BlacklistRegex) error 
	RemoveRegexFromBlackList(regex string) error
}