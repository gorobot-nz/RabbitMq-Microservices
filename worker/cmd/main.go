package main

import "worker/wrappers"

func main() {
	worker := wrappers.NewWorkerWrapper()
	worker.Visit("https://go.dev/learn/")
}
