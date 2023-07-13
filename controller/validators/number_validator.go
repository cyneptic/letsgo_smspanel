package validators

import (
	"github.com/google/uuid"
)

func UUIDValidation(userID string) (bool, uuid.UUID) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return false, uuid.New()
	}
	return true, id
}
