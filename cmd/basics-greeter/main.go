package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/qjcg/go-examples/internal/greet"
)

func Main() int {
	name := flag.String("n", "Jerry", "name of the person to greet")
	flag.Parse()

	greeting := greet.Greet(*name)
	fmt.Println(greeting)

	return 0
}

func main() {
	os.Exit(Main())
}
