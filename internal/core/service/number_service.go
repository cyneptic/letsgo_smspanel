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
		if exists, err := s.db.IsReserved(randomNumber); !exists && err == nil {
			return randomNumber, nil
		}
	}
	return "", errors.New("there is an error in server")

}
func (s *NumberService) BuyNumber(user uuid.UUID, number string) error {
	err := s.db.BuyANumber(user, number)
	if err != nil {
		return err
	}
	return nil
}
func (s *NumberService) SubscribeNumber(user uuid.UUID, number string) error {
	if ok, err := s.db.IsReserved(number); err != nil || !ok {
		return errors.New("user already subscribed to this number")
	}
	err := s.db.SubscribeMe(user, number)
	if err != nil {
		return errors.New("there is an error in subscribing")
	}
	return nil
}
func (s *NumberService) GetSharedNumber() ([]string, error) {
	sharedNumbers, err := s.db.GetSharedANumber()
	if err != nil {
		return []string{}, errors.New("there is no shared number")
	}
	shared := make([]string, 0, len(sharedNumbers))
	for _, number := range sharedNumbers {
		shared = append(shared, number.No)
	}
	return shared, nil
}
