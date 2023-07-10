package validators

import (
	"errors"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func AddWordToBlackListValidator(word entities.BlacklistWord) error {
	if word.Word == "" {
		return errors.New("word feild should not be empty")
	}
	return nil
}
func RemoveWordFromBlackListValidator(word string) error {
	if word == "" {
		return errors.New("word feild should not be empty")
	}
	return nil
}
func AddRegexToBlackListValidator(regex entities.BlacklistRegex) error {
	if regex.Regex == "" {
		return errors.New("regex feild should not be empty")
	}
	return nil
}
func SearchMessageValidator(word string) error {
	if word == "" {
		return errors.New("please provide valid word for search")
	}
	return nil
}