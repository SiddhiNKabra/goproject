package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Qoute struct{
	Quote string
	Author string
}

func main() {
	quotes := []Qoute{}
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
		fmt.Println("Using URL:-", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code:- ", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:- ", err.Error())
	})
	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		qoute := div.Find(".text").Text()
		author := div.Find(".author").Text()
		q := Qoute{
			Quote : qoute,
			Author: author,
		}

		quotes = append(quotes, q)
		// fmt.Printf("Qoute: %s\nBy %s\n\n", qoute, author)
	})
	c.Visit("http://quotes.toscrape.com")

	fmt.Println(quotes)
}
