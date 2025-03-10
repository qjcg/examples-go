//go:build ignore

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	dataTemplate := `// GENERATED CODE, DO NOT EDIT

package generate

const generatedAnswer = {{.Answer}}
`

	tmpl, err := template.New("main").Parse(dataTemplate)
	if err != nil {
		log.Fatal(err)
	}

	data := struct{ Answer int }{
		Answer: 42,
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
