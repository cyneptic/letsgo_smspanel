package ports

import "github.com/google/uuid"

type AdminActionsProviderContract interface {
	DisableUserAccount(targetId uuid.UUID, toggle bool) error
	IsAdmin(userid uuid.UUID) (bool, error)
}

type QueueProviderContract interface {
	SendMessage(sender, msg string, receivers []string) (isSuccessful bool)
}

type MessageProvider interface {
	Publisher(sender, msg string, receivers []string)
}
