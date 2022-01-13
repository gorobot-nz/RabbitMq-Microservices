package wrapper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type CollectorWrapper struct {
	collector *colly.Collector
}

func NewWrapper() *CollectorWrapper {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(e.Request.AbsoluteURL(link))
	})

	return &CollectorWrapper{c}
}

func (w *CollectorWrapper) Run(url string) error {
	err := w.collector.Visit(url)
	if err != nil {
		return err
	}
	return nil
}
