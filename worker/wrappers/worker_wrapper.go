package wrappers

import (
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type WorkerWrapper struct {
	collector *colly.Collector
}

func NewWorkerWrapper() *WorkerWrapper {
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		text := e.Text
		log.Infof("Find link %s", text)
	})

	return &WorkerWrapper{c}
}
