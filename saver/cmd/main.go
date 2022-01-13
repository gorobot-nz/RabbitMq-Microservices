package main

import (
	log "github.com/sirupsen/logrus"
	"saver/wrappers"
)

func main() {
	cfg := wrappers.Config{}

	es := wrappers.NewElasticWrapper(cfg)
	rabbit := wrappers.NewRabbitMQSaverWrapper("amqp://guest:guest@localhost:5672/", es)
	rabbit.Listen()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
