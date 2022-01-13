package wrappers

import (
	"github.com/streadway/amqp"
)

type RabbitMQReceiverWorkerWrapper struct {
	Connection     *amqp.Connection
	ReceiveChannel *amqp.Channel
	ReceiveQueue   *amqp.Queue
	worker         *WorkerWrapper
}

func NewRabbitWorkerMQWrapper(url string, worker *WorkerWrapper) *RabbitMQReceiverWorkerWrapper {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to declare a connection")
	receiveChannel, err := InitChannel(conn)
	failOnError(err, "Failed to declare a receive channel")
	receiveQueue, err := InitQueue(receiveChannel, "tasks")
	failOnError(err, "Failed to declare a receive queue")
	return &RabbitMQReceiverWorkerWrapper{
		conn,
		receiveChannel,
		receiveQueue,
		worker,
	}
}

func (rmq *RabbitMQReceiverWorkerWrapper) Listen(worker *WorkerWrapper) {
	msgs, err := rmq.ReceiveChannel.Consume(
		rmq.ReceiveQueue.Name, // queue
		"",                    // consumer
		true,                  // auto-ack
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                   // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			worker.Visit(string(d.Body))
		}
	}()
}
