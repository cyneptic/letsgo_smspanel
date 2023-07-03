package provider

import (
	"net/http"
	"time"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

const (
	SendSMSProviderHost = "http://localhost:8000"
	httpTimeout         = 5 * time.Second
)

type SnedSMSProviderClient struct {
	client *http.Client
}

func NewSendSMSProviderClient() *SnedSMSProviderClient {
	tr := &http.Transport{}
	cl := &http.Client{
		Transport: tr,
		Timeout:   httpTimeout,
	}

	return &SnedSMSProviderClient{
		client: cl,
	}
}

func (pc *SnedSMSProviderClient) RequestContactList() ([]entities.Contact, error) {
	return nil, nil

}

func (pc *SnedSMSProviderClient) RequestNumber(id string) (entities.Number, error) {
	return entities.Number{}, nil
}
func (pc *SnedSMSProviderClient) RequestUser(id string) (entities.User, error) {
	return entities.User{}, nil
}
