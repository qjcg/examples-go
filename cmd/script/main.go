package main

import (
	"log"

	"github.com/bitfield/script"
)

func main() {
	_, err := script.Get("https://99wttr.in/Montreal?format=3").Stdout()
	if err != nil {
		log.Fatal(err)
	}
}
