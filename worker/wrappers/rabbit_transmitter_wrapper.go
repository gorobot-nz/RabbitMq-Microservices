package wrappers

import "github.com/streadway/amqp"

type RabbitMQTransmitterWorkerWrapper struct {
	Connection      *amqp.Connection
	TransmitChannel *amqp.Channel
	TransmitQueue   *amqp.Queue
}
