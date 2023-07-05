package validators

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

func ValidateNumber(p string) (time.Duration, error) {

	t1, err := strconv.ParseUint(p, 10, 64)
	if err != nil {
		return 0, errors.New("can't convert str to uint")
	}
	// t1 = t1 * 1000000000 * 60 * 60
	t1 = t1 * uint64(time.Hour)
	t := time.Duration(t1)
	return t, nil
}

func ValidatorReciveMessage(msg entities.MessageReciver) (entities.Message, error) {
	var result entities.Message
	if msg.Content == "" {
		return entities.Message{}, errors.New("content can't be empty")

	}
	if msg.Receiver == "" {
		return entities.Message{}, errors.New("receiver can't be empty")

	}
	userid, err := uuid.Parse(msg.UserID)
	if err != nil {
		return entities.Message{}, err
	}
	match, _ := regexp.MatchString("09(1[0-9]|3[1-9]|2[1-9])-?[0-9]{3}-?[0-9]{4}", msg.Sender)
	if !match {
		// panic("incorect phone number")
		return entities.Message{}, fmt.Errorf("can't use %v as correct phone number", msg.Sender)
	}
	result.Content = msg.Content
	result.Receiver = msg.Receiver
	result.Sender = msg.Sender
	result.UserID = userid
	result.CreatedAt = time.Now()

	return result, nil

}
