package validators

import (
	"errors"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func ValidatePhoneBookParam(p entities.PhoneBook) error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
