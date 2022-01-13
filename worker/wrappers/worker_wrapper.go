package wrappers

import (
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type WorkerWrapper struct {
	worker *colly.Collector
	rmq    *RabbitMQTransmitterWorkerWrapper
}

func NewWorkerWrapper(rmq *RabbitMQTransmitterWorkerWrapper) *WorkerWrapper {
	worker := colly.NewCollector()

	return &WorkerWrapper{worker, rmq}
}

func (w *WorkerWrapper) Visit(url string) {
	err := w.worker.Visit(url)

	w.worker.OnHTML("title", func(e *colly.HTMLElement) {
		text := e.Text
		log.Infof("Find info %s", text)
	})

	if err != nil {
		return
	}
}
