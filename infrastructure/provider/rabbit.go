package provider

import (
	"errors"
	"fmt"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

type RabbitQueue struct {
	con         *amqp.Connection
	smsProvider ports.QueueProviderContract
}

func newRabbitConnection() (*amqp.Connection, error) {
	host := os.Getenv("RABBIT_HOST")
	port := os.Getenv("RABBIT_PORT")
	user := os.Getenv("RABBIT_USER")
	password := os.Getenv("RABBIT_PASS")

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, errors.New("error in newRabbitConncetion")
	}
	return connection, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
