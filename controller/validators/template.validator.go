package validators

import (
	"errors"
	"net/url"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func ValidateTemplate(temp entities.Template) error {
	if temp.Content == "" {
		return errors.New("content can not be empty")
	}
	if temp.TempName == "" {
		return errors.New("template name can not be empty")

	}
	return nil
}

func ValidateTempName(p url.Values) error {
	if p.Get("tempName") == "" {
		return errors.New("temp name is empty and can not be empty")
	}
	return nil
}
