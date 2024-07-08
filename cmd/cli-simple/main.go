// A simple cli tool with a small scope of behavior.
package main

import (
	"flag"
	"fmt"
)

// Our custom usage function for CLI help.
func usage() {
	w := flag.CommandLine.Output()
	fmt.Fprintln(w, "Cool Coffee CLI: The best way to get a cup of Joe!")
	fmt.Fprintf(w, "Copyright Cool CLI Corp 1997\n")
	flag.PrintDefaults()
}

func main() {
	nCups := flag.Int("c", 1, "number of cups")
	strength := flag.String("s", "normal", "coffee strength (normal, strong)")
	large := flag.Bool("l", false, "large size coffee")

	flag.Usage = usage // This assignment needs to be BEFORE the parse.

	flag.Parse()

	// Get a slice of all NON-flag arguments.
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "order":
			fmt.Println("PREPARING COFFEE ORDER!")
		case "drink":
			fmt.Println("Drinking coffee. It was GREAT!")
			return
		}
	}

	fmt.Printf("You ordered %d cups of coffee\n", *nCups)

	if *strength == "strong" {
		fmt.Println("You ordered a strong coffee! Buckle up.")
	}

	if *large {
		fmt.Println("You ordered a large coffee! REALLY buckle up.")
	} else {
		fmt.Println("You ordered a regular size coffee.")
	}
}
