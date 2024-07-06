package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed sql
var content embed.FS

func main() {
	demoData, err := content.ReadFile("sql/demo.sql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(demoData))
}
