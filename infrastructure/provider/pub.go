package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (q *RabbitQueue) Publisher(sender, msg string, receivers []string) {
	ch, err := q.con.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"send/sms",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	message := Message{
		Sender:    sender,
		Receivers: receivers,
		Content:   msg,
	}

	marshaledJson, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ch.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshaledJson,
		})
	failOnError(err, "Failed to publish a message")
}
