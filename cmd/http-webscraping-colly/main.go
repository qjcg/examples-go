// A little CLI scraper.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/gocolly/colly/v2"
)

var urls = map[string]string{
	"reddit":     `https://old.reddit.com/r/ThingsCutInHalfPorn/`,
	"metalsucks": `https://www.metalsucks.net/`,
}

var REImgurURL *regexp.Regexp = regexp.MustCompile("i.imgur.com")

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
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		c := colly.NewCollector()
		c = PrintThumbnailURLs(c, os.Stdout)
		err := c.Visit(urls["reddit"])
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	go func() {
		cMetalSucks := colly.NewCollector()
		cMetalSucks = PrintMetalSucksTitles(cMetalSucks, os.Stdout)
		err := cMetalSucks.Visit(urls["metalsucks"])
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
