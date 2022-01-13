package wrappers

import (
	"github.com/gocolly/colly/v2"
)

type CollectorWrapper struct {
	collector *colly.Collector
}

func NewCollectorWrapper() *CollectorWrapper {
	c := colly.NewCollector()
	return &CollectorWrapper{c}
}
