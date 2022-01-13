package main

import (
	log "github.com/sirupsen/logrus"
	"worker/wrappers"
)

func main() {
	rabbitTransmitter := wrappers.NewRabbitMQTransmitterWorkerWrapper("amqp://guest:guest@localhost:5672/")
	worker := wrappers.NewWorkerWrapper(rabbitTransmitter)
	rabbitReceiver := wrappers.NewRabbitMQReceiverWorkerWrapper("amqp://guest:guest@localhost:5672/", worker)
	rabbitReceiver.Listen()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
