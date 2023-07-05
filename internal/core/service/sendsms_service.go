package service

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/infrastructure/provider"
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
)

type SendSMSService struct {
	db ports.SnedSMSRepositoryContract
	pv ports.QueueContract
}

func NewSendSMSService() *SendSMSService {
	db := repositories.NewGormDatabase()
	pv := provider.NewQueueConnection()
	return &SendSMSService{
		db: db,
		pv: pv,
	}
}

// !SendToContactList
func (svc *SendSMSService) SendToContactList(msg entities.Message) error {
	dataContacts, err := svc.db.RequestContactList()
	dataColloctionNumber := make([]string, 0)
	if err != nil {
		panic(err)
	}
	for _, contact := range dataContacts {
		dataColloctionNumber = append(dataColloctionNumber, contact.PhoneNumber)
	}
	svc.pv.Publisher(msg.Sender, msg.Content, dataColloctionNumber)
	return nil
}

// !SendToNumber
func (svc *SendSMSService) SendToNumber(msg entities.Message) error {
	dataPhone, err := svc.db.RequestNumber(msg.UserID)
	dataColloctionNumber := make([]string, 0)
	if err != nil {
		panic(err)
	}
	dataColloctionNumber = append(dataColloctionNumber, dataPhone.No)
	svc.pv.Publisher(msg.Sender, msg.Content, dataColloctionNumber)
	return nil
}

// !SendToUserName
func (svc *SendSMSService) SendToUser(msg entities.Message) error {
	dataUser, err := svc.db.RequestUser(msg.UserID)
	dataColloctionNumber := make([]string, 0)
	if err != nil {
		panic(err)
	}
	dataColloctionNumber = append(dataColloctionNumber, dataUser.PhoneNumber)
	svc.pv.Publisher(msg.Sender, msg.Content, dataColloctionNumber)
	return nil
}

// !SendToContactListInterval
func (svc *SendSMSService) SendToContactListInterval(msg entities.Message, interval time.Duration) error {
	dataContacts, err := svc.db.RequestContactList()

	dataColloctionNumber := make([]string, 0)
	if err != nil {
		panic(err)
	}
	for _, contact := range dataContacts {
		dataColloctionNumber = append(dataColloctionNumber, contact.PhoneNumber)
	}

	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				svc.pv.Publisher(msg.Sender, msg.Content, dataColloctionNumber)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return nil
}
