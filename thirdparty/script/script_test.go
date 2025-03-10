package script

import (
	"fmt"
	"log"
	"strings"

	"github.com/bitfield/script"
)

func ExampleFile() {
	lineCount, err := script.File("./testdata/passwd").Match("nobody").CountLines()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lineCount)

	// Output: 1
}

func Example_newPipeWithReader() {
	r := strings.NewReader(`first line
second line
third line
`)

	lineCount, err := script.NewPipe().WithReader(r).CountLines()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lineCount)

	// Output: 3
}
