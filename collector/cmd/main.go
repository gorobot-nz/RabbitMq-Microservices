package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(e.Request.AbsoluteURL(link))
	})

	err := c.Visit("https://go.dev/learn/")
	if err != nil {
		return
	}
}
