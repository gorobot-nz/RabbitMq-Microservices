package wrappers

import (
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type WorkerWrapper struct {
	worker *colly.Collector
}

func NewWorkerWrapper(rmq *RabbitMQWorkerWrapper) *WorkerWrapper {
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		text := e.Text
		log.Infof("Find info %s", text)
		rmq.Send(text)
	})

	return &WorkerWrapper{c}
}

func (w *WorkerWrapper) Visit(url string) {
	err := w.worker.Visit(url)
	if err != nil {
		return
	}
}
