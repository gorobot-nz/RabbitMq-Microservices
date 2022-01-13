package wrappers

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQWrapper struct {
	Connection      *amqp.Connection
	ReceiveChannel  *amqp.Channel
	TransmitChannel *amqp.Channel
	ReceiveQueue    *amqp.Queue
	TransmitQueue   *amqp.Queue
}

func NewRabbitMQWrapper(url string) *RabbitMQWrapper {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	receiveChannel, err := InitChannel(conn)
	transmitChannel, err := InitChannel(conn)
	receiveQueue, err := InitQueue(receiveChannel, "tasks")
	transmitQueue, err := InitQueue(receiveChannel, "tasks_result")
	return &RabbitMQWrapper{
		conn,
		receiveChannel,
		transmitChannel,
		&receiveQueue,
		&transmitQueue}
}

func InitChannel(connection *amqp.Connection) (*amqp.Channel, error) {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return ch, nil
}

func InitQueue(chanel *amqp.Channel, name string) (amqp.Queue, error) {
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
	return q, nil
}
