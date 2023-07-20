package main

import (
	"flag"
	"fmt"

	"github.com/qjcg/go-examples/internal/greet"
)

func main() {
	name := flag.String("n", "Jerry", "name of the person to greet")
	flag.Parse()

	greeting := greet.Greet(*name)
	fmt.Println(greeting)
}
