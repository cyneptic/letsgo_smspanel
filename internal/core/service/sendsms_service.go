package service

import (
	"time"

	"github.com/cyneptic/letsgo-smspanel/infrastructure/provider"
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/google/uuid"
)

type SendSMSService struct {
	db  ports.SnedSMSRepositoryContract
	pv  ports.MessageProvider
	ich chan IntervalMessage
}

type IntervalMessage struct {
	sender, content string
	receivers       []string
}

func NewSendSMSService() *SendSMSService {
	db := repositories.NewGormDatabase()
	pv, err := provider.NewQueueConnection()
	if err != nil {
		panic(err)
	}
	ich := make(chan IntervalMessage)
	result := &SendSMSService{
		db:  db,
		pv:  pv,
		ich: ich,
	}
	go result.SendInterval()
	return result
}

// !SendToContactList
func (svc *SendSMSService) SendToContactList(msg entities.Message) error {
	dataContacts, err := svc.db.RequestContactList(msg.UserID)
	dataColloctionNumber := make([]string, 0)
	if err != nil {
		panic(err)
	}
	for _, contact := range dataContacts {
		dataColloctionNumber = append(dataColloctionNumber, contact.PhoneNumber)
	}

	price, err := svc.db.GetGroupPrice()
	if err != nil {
		return err
	}
	err = svc.CollectCost(msg.UserID, price)
	if err != nil {
		return err
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

	price, err := svc.db.GetSinglePrice()
	if err != nil {
		return err
	}
	err = svc.CollectCost(msg.UserID, price)
	if err != nil {
		return err
	}

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

	price, err := svc.db.GetSinglePrice()
	if err != nil {
		return err
	}
	err = svc.CollectCost(msg.UserID, price)
	if err != nil {
		return err
	}

	svc.pv.Publisher(msg.Sender, msg.Content, dataColloctionNumber)
	return nil
}

func (svc *SendSMSService) CollectCost(userid uuid.UUID, amount int) error {
	err := svc.db.WithdrawFromWallet(userid, amount)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SendSMSService) SendInterval() {
	for {
		select {
		case msg := <-svc.ich:
			svc.pv.Publisher(msg.sender, msg.content, msg.receivers)
		}
	}
}

// !SendToContactListInterval
func (svc *SendSMSService) SendToContactListInterval(msg entities.Message, interval time.Duration) error {
	dataContacts, err := svc.db.RequestContactList(msg.UserID)

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
				svc.ich <- IntervalMessage{msg.Sender, msg.Content, dataColloctionNumber}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return nil
}
