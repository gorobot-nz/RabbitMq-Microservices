package wrappers

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQWrapper struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Query      string
}

func NewRabbitMQWrapper(url string) *RabbitMQWrapper {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	q, err := ch.QueueDeclare(
		"tasks",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return &RabbitMQWrapper{conn, ch, q.Name}
}
