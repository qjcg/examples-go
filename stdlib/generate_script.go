//go:build ignore

package main

import (
	"log"
	"math/rand/v2"
	"os"
	"text/template"
)

func main() {
	dataTemplate := `// GENERATED CODE, DO NOT EDIT

package stdlib

const generatedAnswer = {{.Answer}}
`

	// Random source is seeded with static values, so will always
	// produce the same results. This is what we want for testing.
	r := rand.New(rand.NewPCG(1, 2))

	tmpl, err := template.New("main").Parse(dataTemplate)
	if err != nil {
		log.Fatal(err)
	}

	data := struct{ Answer int }{
		Answer: r.IntN(50),
	}

	f, err := os.Create("generate_data.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}
