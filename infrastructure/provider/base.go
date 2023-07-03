package provider

import (
	"os"
)

type Message struct {
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
	Content   string   `json:"content"`
}

func NewSMSProvider() *KavenegarProvider {
	token := os.Getenv("SMS_PROVIDER_TOKEN")
	return &KavenegarProvider{
		token:     token,
		urlFormat: "https://api.kavenegar.com/v1/%s/sms/send.json?receptor=%s&sender=%s&message=%s",
	}
}

func NewQueueConnection() (*RabbitQueue, error) {
	conn, err := newRabbitConnection()
	if err != nil {
		return nil, err
	}
	return &RabbitQueue{
		con:         conn,
		smsProvider: NewSMSProvider(),
	}, nil
}
