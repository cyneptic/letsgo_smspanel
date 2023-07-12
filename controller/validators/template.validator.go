package validators

import (
	"errors"

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
