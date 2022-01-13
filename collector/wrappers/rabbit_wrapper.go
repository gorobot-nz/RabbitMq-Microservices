package wrappers

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQCollectorWrapper struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      *amqp.Queue
}

func NewRabbitMQCollectorWrapper(url string) *RabbitMQCollectorWrapper {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to declare a connection")
	ch, err := InitChannel(conn)
	failOnError(err, "Failed to declare a channel")
	q, err := InitQueue(ch, "tasks")
	failOnError(err, "Failed to declare a queue")
	return &RabbitMQCollectorWrapper{conn, ch, q}
}

func (rmq *RabbitMQCollectorWrapper) Send(message string) {
	err := rmq.Channel.Publish(
		"",
		rmq.Queue.Name,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatal(err)
	}
}
