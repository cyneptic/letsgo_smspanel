package ports

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

type SendSMSServiceContract interface {
	SendToContactList(contactList []entities.Contact, msg string) error
	SendToNumber(contactList entities.Number, msg string) error
	SendToUserName(contactList entities.User, msg string) error
	SendToContactListInterval(contactList []entities.Contact, msg string, interval time.Time) error
}
