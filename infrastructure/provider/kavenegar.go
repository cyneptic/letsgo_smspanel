package provider

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"strconv"
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
	isDebug, err := strconv.ParseBool(
		os.Getenv("DEBUG"),
	)
	if err != nil {
		return false
	}
	if isDebug {
		for _, receiver := range receivers {
			log.Info(fmt.Sprintf("*******\nsuccess\nsender: %s\nreceiver: %s\nmessage: %s", sender, receiver, msg))
		}
		return true
	} else {
		receiverNumbers := s.makeReceivers(receivers)
		response, _ := http.Get(s.makeRequestUrl(sender, receiverNumbers, msg))
		return response.StatusCode == http.StatusOK
	}

}
