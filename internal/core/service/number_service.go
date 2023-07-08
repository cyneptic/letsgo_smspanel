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
	db ports.NumberServiceContract
}

func NewNumberService() *NumberService {
	db := repositories.NewGormDatabase()
	return &NumberService{
		db: db,
	}
}

func (s *NumberService) GenerateNumber() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := "1000" + fmt.Sprintf("%07d", rand.Intn(10000000))
	return randomNumber
}
func (s *NumberService) BuyNumber(user string, number string) error {
	userID, _ := uuid.Parse(user)
	numberID, _ := uuid.Parse(number)
	if ok, err := s.db.IsNumberFree(number); err != nil || !ok {
		return err
	}
	err := s.db.BuyANumber(userID, numberID)
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
