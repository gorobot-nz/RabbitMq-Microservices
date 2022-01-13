package main

import (
	"collector/wrapper"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	collector := wrapper.NewWrapper()
	err := collector.Run("https://go.dev/learn/")
	if err != nil {
		return
	}
}
