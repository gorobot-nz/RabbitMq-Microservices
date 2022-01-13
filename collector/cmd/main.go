package main

import (
	"collector/wrappers"
)

func main() {
	rabbitmq := wrappers.NewRabbitMQWrapper("amqp://guest:guest@localhost:5672/")
	collector := wrappers.NewCollectorWrapper(rabbitmq)
	err := collector.Run("https://go.dev/learn/")
	if err != nil {
		return
	}
}
