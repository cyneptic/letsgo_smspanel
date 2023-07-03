package ports

import "github.com/cyneptic/letsgo-smspanel/internal/core/entities"

type SnedSMSRepositoryContract interface {
	RequestContactList() ([]entities.Contact, error)
	RequestNumber(id string) (entities.Number, error)
	RequestUser(id string) (entities.User, error)
}
