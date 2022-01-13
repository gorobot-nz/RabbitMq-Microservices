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
	collector := wrappers.NewCollectorWrapper()
	err := collector.Run("https://go.dev/learn/")
	if err != nil {
		return
	}
}
