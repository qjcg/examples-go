package stdlib

import "fmt"

//go:generate go run generate_script.go

func ExampleGoGenerate() {
	fmt.Println(generatedAnswer)

	// Output: 38
}
