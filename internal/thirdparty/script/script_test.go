package script

import (
	"log"
	"time"

	"github.com/bitfield/script"
)

func ExampleGet() {
	_, err := script.Get("https://httpbingo.org/stream/3").JQ(".id").Stdout()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)
	// Output:
	// 0
	// 1
	// 2
}

func ExampleFile() {
	_, err := script.File("/etc/passwd").Match("root")
	if err != nil {
		log.Fatal(err)
	}

	// Output:b
}
