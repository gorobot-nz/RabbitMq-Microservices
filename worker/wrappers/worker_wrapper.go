package wrappers

import (
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type CollectorWrapper struct {
	collector *colly.Collector
}

func NewCollectorWrapper() *CollectorWrapper {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		log.Infof("Find link %s", e.Request.AbsoluteURL(link))
	})

	return &CollectorWrapper{c}
}
