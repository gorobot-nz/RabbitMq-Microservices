package wrappers

import "github.com/streadway/amqp"

type RabbitMQSaverWrapper struct {
	Connection     *amqp.Connection
	ReceiveChannel *amqp.Channel
	ReceiveQueue   *amqp.Queue
	es             *ElasticWrapper
}

func NewRabbitMQSaverWrapper(url string, es *ElasticWrapper) *RabbitMQSaverWrapper {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to declare a connection")
	receiveChannel, err := InitChannel(conn)
	failOnError(err, "Failed to declare a receive channel")
	receiveQueue, err := InitQueue(receiveChannel, "tasks")
	failOnError(err, "Failed to declare a receive queue")
	return &RabbitMQSaverWrapper{
		conn,
		receiveChannel,
		receiveQueue,
		es,
	}
}

func (rmq *RabbitMQSaverWrapper) Listen() {
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
			rmq.es.Save(string(d.Body), "")
		}
	}()
}
