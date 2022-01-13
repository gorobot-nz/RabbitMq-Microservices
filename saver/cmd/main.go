package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"saver/wrappers"
	"strings"
)

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func CheckEnvVars() {
	requiredEnvs := []string{"ELASTIC_HOST", "ELASTIC_USER", "ELASTIC_PASSWORD", "RABBIT_HOST"}
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
	err := InitConfig()
	if err != nil {
		log.Fatal("Config initial error")
	}
	CheckEnvVars()

	config := wrappers.Config{
		Host:     os.Getenv("ELASTIC_HOST"),
		Username: os.Getenv("ELASTIC_USER"),
		Password: os.Getenv("ELASTIC_PASSWORD"),
		Index:    viper.GetString("index"),
	}
	es := wrappers.NewElasticWrapper(config)
	rabbit := wrappers.NewRabbitMQSaverWrapper("amqp://guest:guest@localhost:5672/", es)
	rabbit.Listen()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
