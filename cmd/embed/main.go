package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed testdata/greeting.txt
var greeting string

//go:embed testdata/sql
var content embed.FS

func main() {
	demoData, err := content.ReadFile("testdata/sql/demo.sql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(demoData))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Greeting: %s\n", greeting)
}
