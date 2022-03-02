package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("spa1.scrape.center"),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	extensions.RandomUserAgent(c)
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println("html:", e)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		fmt.Println("request:", r)
	})

	c.OnResponse(func(r *colly.Response) {

		fmt.Println("Body:", string(r.Body))
	})

	c.Visit("https://spa1.scrape.center")
}
