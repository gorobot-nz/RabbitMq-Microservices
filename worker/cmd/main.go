package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"worker/wrappers"
)

func CheckEnvVars() {
	requiredEnvs := []string{"RABBIT_HOST"}
	var msg []string
	for _, el := range requiredEnvs {
		val, exists := os.LookupEnv(el)
		if !exists || len(val) == 0 {
			msg = append(msg, el)
		}
	}
	if len(msg) > 0 {
		log.Fatal(strings.Join(msg, ", "), " env(s) not set")
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env error: %s", err.Error())
	}

	CheckEnvVars()

	rabbitTransmitter := wrappers.NewRabbitMQTransmitterWorkerWrapper(os.Getenv("RABBIT_HOST"))
	worker := wrappers.NewWorkerWrapper(rabbitTransmitter)
	rabbitReceiver := wrappers.NewRabbitMQReceiverWorkerWrapper(os.Getenv("RABBIT_HOST"), worker)
	rabbitReceiver.Listen()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
