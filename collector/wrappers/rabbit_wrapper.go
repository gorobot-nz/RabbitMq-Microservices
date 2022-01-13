package wrappers

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQWrapper struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      *amqp.Queue
}

func NewRabbitMQWrapper(url string) *RabbitMQWrapper {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to declare a connection")
	ch, err := InitChannel(conn)
	failOnError(err, "Failed to declare a channel")
	q, err := InitQueue(ch, "tasks")
	failOnError(err, "Failed to declare a queue")
	return &RabbitMQWrapper{conn, ch, q}
}

func (rmq *RabbitMQWrapper) Send(message string) {
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InitChannel(connection *amqp.Connection) (*amqp.Channel, error) {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return ch, nil
}

func InitQueue(chanel *amqp.Channel, name string) (*amqp.Queue, error) {
	q, err := chanel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return &q, nil
}
