package ports

import "github.com/google/uuid"

type AdminActionsProviderContract interface {
	DisableUserAccount(userId uuid.UUID) error
}
