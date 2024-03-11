//go:build ignore

package main

import (
	"embed"
	"io"
	"log"
	"os"
	"text/template"
)

//go:embed data.go.tmpl
var dataTemplate embed.FS

type App struct {
	w    io.Writer
	tmpl *template.Template
	data any
}

func newApp(w io.Writer, data any) (*App, error) {
	var app App
	var err error

	tmpl, err := template.ParseFS(dataTemplate, "data.go.tmpl")
	if err != nil {
		return &app, err
	}

	app.w = w
	app.data = data
	app.tmpl = tmpl

	return &app, err
}

func (app *App) Execute() error {
	return app.tmpl.Execute(app.w, app.data)
}

func main() {
	data := map[string]any{
		"items": map[string]any{
			"foo": 42,
			"bar": 99,
			"baz": 102,
		},
	}

	f, err := os.Create("data.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	app, err := newApp(f, data)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
