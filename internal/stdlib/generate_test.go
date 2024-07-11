package stdlib

import "fmt"

//go:generate go run generate_script.go

func ExampleGoGenerate() {
	fmt.Printf("generatedAnswer: %d\n", generatedAnswer)

	// Output:
	// generatedAnswer: 38
}
