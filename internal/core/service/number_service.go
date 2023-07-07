package service

import (
	"fmt"
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
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
func (s *NumberService) BuyNumber(user string, number string) {

}
func (s *NumberService) SubscribeNumber(user, number string) error {
	return nil
}
func (s *NumberService) GetSharedNumber() string {
	return ""
}
