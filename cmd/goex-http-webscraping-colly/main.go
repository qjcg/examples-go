// A little CLI scraper for the ThingsCutInHalf reddit.
// https://www.reddit.com/r/ThingsCutInHalfPorn/
package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

const RedditURL = `https://old.reddit.com/r/ThingsCutInHalfPorn/`

var REValidURL *regexp.Regexp = regexp.MustCompile("imgur")

func main() {
	c := colly.NewCollector()

	c.OnHTML("a.thumbnail", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if REValidURL.MatchString(url) {
			fmt.Println(url)
		}
	})

	c.Visit(RedditURL)
}
