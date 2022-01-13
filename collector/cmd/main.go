package main

import (
	"collector/wrappers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
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

	rabbitmq := wrappers.NewRabbitMQCollectorWrapper(os.Getenv("RABBIT_HOST"))
	collector := wrappers.NewCollectorWrapper(rabbitmq)
	err := collector.Run("https://go.dev/learn/")
	if err != nil {
		return
	}
}
