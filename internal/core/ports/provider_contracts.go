package ports

import "github.com/google/uuid"

type AdminActionsProviderContract interface {
	DisableUserAccount(targetId uuid.UUID, toggle bool) error
	IsAdmin(userid uuid.UUID) (bool, error)
}
