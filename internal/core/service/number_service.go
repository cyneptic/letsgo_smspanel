package service

import (
	"errors"
	"fmt"
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type NumberService struct {
	db ports.NumberRepositoryContract
}

func NewNumberService() *NumberService {
	db := repositories.NewGormDatabase()
	return &NumberService{
		db: db,
	}
}

func (s *NumberService) GenerateNumber() (string, error) {
	rand.Seed(time.Now().UnixNano())
	numberPrefix := []string{
		"1000",
		"2000",
		"3000",
		"4000",
		"5000",
	}
	for i := 0; i < 10; i++ {
		randomNumber := numberPrefix[rand.Intn(len(numberPrefix))] + fmt.Sprintf("%07d", rand.Intn(10000000))
		if !s.db.IsReserved(randomNumber) {
			return randomNumber, nil
		}
	}
	return "", errors.New("there is an error in server")

}
func (s *NumberService) BuyNumber(user string) error {
	userID, _ := uuid.Parse(user)
	number, err := s.GenerateNumber()
	if err != nil {
		return err
	}
	err = s.db.BuyANumber(userID, number)
	if err != nil {
		return err
	}
	return nil
}
func (s *NumberService) SubscribeNumber(user, number string) error {
	if ok, err := s.db.IsSubscribable(user, number); err != nil || !ok {
		return err
	}
	s.db.SubscribeMe(user, number)
	return nil
}
func (s *NumberService) GetSharedNumber() (string, error) {
	number, err := s.db.GetShareANumber()
	if err != nil {
		return "", errors.New("there is no shared number")
	}
	return number, nil
}
