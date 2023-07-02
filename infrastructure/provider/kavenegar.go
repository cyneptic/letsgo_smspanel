package provider

import (
	"fmt"
	"strings"
)

type KavenegarProvider struct {
	urlFormat string
	token     string
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
	//response, _ := http.Get(s.makeRequestUrl(sender, s.makeReceivers(receivers), msg))
	//if response.StatusCode == http.StatusOK {
	//	return true
	//}
	return true
}
