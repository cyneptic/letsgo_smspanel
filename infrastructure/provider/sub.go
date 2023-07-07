package provider

import (
	"encoding/json"
	"fmt"
)

func (q *RabbitQueue) StartConsuming() {
	ch, err := q.con.Channel()
	queue, err := ch.QueueDeclare(
		"send/sms",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	messages, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for msg := range messages {
			var message Message
			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				fmt.Println("there is error in UnMarshaling")
				continue
			}
			result := q.smsProvider.SendMessage(message.Sender, message.Content, message.Receivers)
			if result {
				fmt.Println("successful")
			}
		}
	}()

	forever := make(chan bool)
	<-forever
}
