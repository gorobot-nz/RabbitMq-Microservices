package wrappers

import "github.com/streadway/amqp"

type RabbitMQTransmitterWorkerWrapper struct {
	Connection      *amqp.Connection
	TransmitChannel *amqp.Channel
	TransmitQueue   *amqp.Queue
}

func NewRabbitMQTransmitterWorkerWrapper(url string) *RabbitMQTransmitterWorkerWrapper {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to declare a connection")
	transmitChannel, err := InitChannel(conn)
	failOnError(err, "Failed to declare a receive channel")
	transmitQueue, err := InitQueue(transmitChannel, "tasks_results")
	failOnError(err, "Failed to declare a receive queue")
	return &RabbitMQTransmitterWorkerWrapper{
		conn,
		transmitChannel,
		transmitQueue,
	}
}

func (rmq *RabbitMQTransmitterWorkerWrapper) Send(message string) {

}
