package main

import (
	"collector/wrappers"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	rabbitmq := wrappers.NewRabbitMQWrapper("amqp://guest:guest@localhost:5672/")
	collector := wrappers.NewCollectorWrapper(rabbitmq)
	err := collector.Run("https://go.dev/learn/")
	if err != nil {
		return
	}
}
