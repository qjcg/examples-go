package main

import "fmt"

//go:generate go run gen.go

func main() {
	fmt.Printf("foo bar baz -> %d %d %d\n", foo, bar, baz)
}
