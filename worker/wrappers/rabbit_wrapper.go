package wrappers

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQWorkerWrapper struct {
	Connection      *amqp.Connection
	ReceiveChannel  *amqp.Channel
	TransmitChannel *amqp.Channel
	ReceiveQueue    *amqp.Queue
	TransmitQueue   *amqp.Queue
}

func NewRabbitWorkerMQWrapper(url string) *RabbitMQWorkerWrapper {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to declare a connection")
	receiveChannel, err := InitChannel(conn)
	failOnError(err, "Failed to declare a receive channel")
	transmitChannel, err := InitChannel(conn)
	failOnError(err, "Failed to declare a transmit channel")
	receiveQueue, err := InitQueue(receiveChannel, "tasks")
	failOnError(err, "Failed to declare a receive queue")
	transmitQueue, err := InitQueue(receiveChannel, "tasks_result")
	failOnError(err, "Failed to declare a transmit queue")
	return &RabbitMQWorkerWrapper{
		conn,
		receiveChannel,
		transmitChannel,
		receiveQueue,
		transmitQueue}
}

func (rmq *RabbitMQWorkerWrapper) Listen(worker *WorkerWrapper) {
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

func (rmq *RabbitMQWorkerWrapper) Send(message string) {
	err := rmq.TransmitChannel.Publish(
		"",
		rmq.TransmitQueue.Name,
		false,
		false,
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
