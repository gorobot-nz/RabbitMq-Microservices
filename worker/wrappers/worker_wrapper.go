package wrappers

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type WorkerWrapper struct {
	worker *colly.Collector
}

func NewWorkerWrapper(rmq *RabbitMQTransmitterWorkerWrapper) *WorkerWrapper {
	worker := colly.NewCollector()

	worker.OnHTML("title", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		title := e.Text
		task := CreateMessage(url, title)
		log.Infof("Find info %s", task)
		rmq.Send(task)
	})

	return &WorkerWrapper{worker}
}

func (w *WorkerWrapper) Visit(url string) {
	err := w.worker.Visit(url)
	if err != nil {
		return
	}
}

func CreateMessage(url, title string) string {
	message := fmt.Sprintf("URL: %s\nTITLE: %s", url, title)
	return message
}
