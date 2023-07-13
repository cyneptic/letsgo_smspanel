package provider

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type KavenegarProvider struct {
	urlFormat string
	token     string
}

func NewSMSProvider() *KavenegarProvider {
	token := os.Getenv("SMS_PROVIDER_TOKEN")
	return &KavenegarProvider{
		token:     token,
		urlFormat: "https://api.kavenegar.com/v1/%s/sms/send.json?receptor=%s&sender=%s&message=%s",
	}
}

func (s *KavenegarProvider) makeRequestUrl(sender, receiver, message string) string {
	url := fmt.Sprintf(s.urlFormat, s.token, receiver, sender, message)
	return url

}

func (s *KavenegarProvider) makeReceivers(receiver []string) string {
	strReceivers := strings.Join(receiver, ",")
	return strReceivers

}

func (s *KavenegarProvider) SendMessage(sender, msg string, receivers []string) (isSuccessful bool) {

	receiverNumbers := s.makeReceivers(receivers)

	response, _ := http.Get(s.makeRequestUrl(sender, receiverNumbers, msg))

	return response.StatusCode == http.StatusOK
}
