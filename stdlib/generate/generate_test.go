package generate

import "fmt"

//go:generate go run generate_script.go

func Example_goGenerate() {
	fmt.Println(generatedAnswer)

	// Output: 42
}
