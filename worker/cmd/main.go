package main

import "worker/wrappers"

func main() {
	rabbitmq := wrappers.NewRabbitWorkerMQWrapper("amqp://guest:guest@localhost:5672/")
	worker := wrappers.NewWorkerWrapper(rabbitmq)
	worker.Visit("https://go.dev/learn/")
}
