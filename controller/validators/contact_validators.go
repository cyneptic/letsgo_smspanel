package validators

import (
	"errors"
	"regexp"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func ValidatePhoneNumber(no string) error {
	if no == "" {
		return errors.New("PhoneNumber must be filled")
	}

	iranianMobilePattern := regexp.MustCompile(`^09\d{9}$`)

	if !iranianMobilePattern.MatchString(no) {
		return errors.New("invalid phone number format")
	}

	return nil
}

func ValidateUsername(c entities.Contact) error {
	if c.Username == "" {
		return errors.New("username must be filled")

	}
	return nil
}

func ValidateContactParam(c entities.Contact) error {
	return ValidatePhoneNumber(c.PhoneNumber)
}

func ValidateUpdateContactParam(c entities.Contact) error {
	if c.FirstName == "" && c.LastName == "" && c.Username == "" && c.PhoneNumber == "" {
		return errors.New("at least one field must be filled")
	}

	if c.PhoneNumber != "" {
		if err := ValidatePhoneNumber(c.PhoneNumber); err != nil {
			return err
		}
	}

	return nil
}
