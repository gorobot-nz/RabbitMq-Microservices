package wrappers

import (
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type CollectorWrapper struct {
	collector *colly.Collector
	rmq       *RabbitMQWrapper
}

func NewCollectorWrapper(rmq *RabbitMQWrapper) *CollectorWrapper {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		log.Infof("Find link %s", e.Request.AbsoluteURL(link))
		rmq.Send(e.Request.AbsoluteURL(link))
	})

	return &CollectorWrapper{c, rmq}
}

func (w *CollectorWrapper) Run(url string) error {
	err := w.collector.Visit(url)
	if err != nil {
		return err
	}
	return nil
}
