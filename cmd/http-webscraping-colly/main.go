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

var (
	urls = map[string]string{
		"reddit":     `https://old.reddit.com/r/ThingsCutInHalfPorn/`,
		"metalsucks": `https://www.metalsucks.net/`,
	}
)

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

func PrintMetalSucksTitles(c *colly.Collector, w io.Writer) *colly.Collector {
	c.OnHTML(".post-title", func(e *colly.HTMLElement) {
		fmt.Fprintln(w, e.Text)
	})

	return c
}

func main() {
	c := colly.NewCollector()
	c = PrintThumbnailURLs(c, os.Stdout)
	c.Visit(urls["reddit"])

	cMetalSucks := colly.NewCollector()
	cMetalSucks = PrintMetalSucksTitles(c, os.Stdout)
	cMetalSucks.Visit(urls["metalsucks"])
}
