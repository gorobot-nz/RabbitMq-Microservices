package wrappers

import (
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type CollectorWrapper struct {
	collector *colly.Collector
	rmq       *RabbitMQCollectorWrapper
}

func NewCollectorWrapper(rmq *RabbitMQCollectorWrapper) *CollectorWrapper {
	collector := colly.NewCollector()

	return &CollectorWrapper{collector, rmq}
}

func (w *CollectorWrapper) Run(url string) error {
	w.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		log.Infof("Find link %s", e.Request.AbsoluteURL(link))
		w.rmq.Send(e.Request.AbsoluteURL(link))
	})
	err := w.collector.Visit(url)
	if err != nil {
		return err
	}
	return nil
}
