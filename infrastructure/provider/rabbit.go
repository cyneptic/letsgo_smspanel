package provider

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitQueue struct {
	con         *amqp.Connection
	smsProvider ports.QueueProviderContract
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

func newRabbitConnection() (*amqp.Connection, error) {
	host := os.Getenv("RABBIT_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")

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
