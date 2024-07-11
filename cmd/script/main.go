// Example usage of bitfield/script to print weather from wttr.in.
//
// For wttr.in formatting options, see:
// https://github.com/chubin/wttr.in?tab=readme-ov-file#one-line-output
package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/bitfield/script"
)

func main() {
	flagCity := flag.String("c", "montreal", "City")
	flagFormat := flag.String("f", "%l: %c %t %C", "Weather format string")
	flag.Parse()

	weatherURL := fmt.Sprintf("https://wttr.in/%s?format=%s", url.PathEscape(*flagCity), url.QueryEscape(*flagFormat))
	_, err := script.Get(weatherURL).Stdout()
	if err != nil {
		log.Fatal(err)
	}
}
