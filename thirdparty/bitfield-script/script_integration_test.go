//go:build integration

package script

import (
	"log"

	"github.com/bitfield/script"
)

func ExampleGet() {
	_, err := script.Get("https://httpbingo.org/json").JQ(".slideshow.slides[0].title").Stdout()
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// "Wake up to WonderWidgets!"
}
