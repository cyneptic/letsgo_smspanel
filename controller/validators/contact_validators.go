package validators

import (
	"errors"
	"regexp"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func ValidateContactByUsernameParam(c entities.Contact) error {
	if c.FirstName == "" {
		return errors.New("first name must be filled")
	}
	err := ValidateUsername(c.Username)
	if err != nil {
		return err
	}
	err = ValidatePhoneNumber(c.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func ValidateContact(c entities.Contact) error {
	if c.FirstName == "" {
		return errors.New("first name must be filled")
	}

	if c.Username != "" {
		err := ValidateUsername(c.Username)
		if err != nil {
			return err
		}
	}

	err := ValidatePhoneNumber(c.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func ValidateUsername(username string) error {
	if username == "" {
		return errors.New("username must be filled")
	}
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if len(username) > 20 {
		return errors.New("username must not exceed 20 characters")
	}
	if !isValidUsername(username) {
		return errors.New("username contains invalid characters")
	}
	return nil
}

func isValidUsername(username string) bool {
	match, _ := regexp.MatchString("^[A-Za-z0-9_]+$", username)
	return match
}

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

func ValidateUpdateContactParam(c entities.Contact) error {
	if c.FirstName == "" && c.LastName == "" && c.Username == "" && c.PhoneNumber == "" {
		return errors.New("at least one field must be filled")
	}

	if c.Username != "" {
		if err := ValidateUsername(c.Username); err != nil {
			return err
		}
	}

	if c.PhoneNumber != "" {
		if err := ValidatePhoneNumber(c.PhoneNumber); err != nil {
			return err
		}
	}

	return nil
}
