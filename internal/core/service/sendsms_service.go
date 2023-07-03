package service

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
)

type SendSMSService struct {
	db ports.SnedSMSRepositoryContract
	pv ports.QueueContract
}

func NewSendSMSService() *SendSMSService {
	return &SendSMSService{
		pv: pv,
	}
}

func (svc *SendSMSService) SendToContactList(contactList []entities.Contact, msg string) error {
	return nil
}
func (svc *SendSMSService) SendToNumber(contactList entities.Number, msg string) error {
	return nil
}
func (svc *SendSMSService) SendToUserName(contactList entities.User, msg string) error {
	return nil
}
func (svc *SendSMSService) SendToContactListInterval(contactList []entities.Contact, msg string, interval time.Time) error {
	return nil
}
