// A little CLI scraper for the ThingsCutInHalf reddit.
// https://www.reddit.com/r/ThingsCutInHalfPorn/
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/gocolly/colly/v2"
)

const RedditURL = `https://old.reddit.com/r/ThingsCutInHalfPorn/`

var REImgurURL *regexp.Regexp = regexp.MustCompile("imgur")

func PrintThumbnailURLs(c *colly.Collector, w io.Writer) *colly.Collector {
	c.OnHTML("a.thumbnail", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if REImgurURL.MatchString(url) {
			fmt.Fprintln(w, url)
		}
	})

	return c
}

func main() {
	c := colly.NewCollector()
	c = PrintThumbnailURLs(c, os.Stdout)
	c.Visit(RedditURL)
}
