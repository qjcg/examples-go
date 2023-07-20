// A simple demo web application.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

type GreetResponse struct {
	Greeting string `json:"greeting"`
}

func Index(w http.ResponseWriter, req *http.Request) {
	render.JSON(w, req, "Hello CHI!")
}

func Greet(w http.ResponseWriter, req *http.Request) {
	var gr GreetResponse
	switch chi.URLParam(req, "lang") {
	case "es":
		gr.Greeting = "Hola!"
	case "fr":
		gr.Greeting = "Allo!"
	case "en":
		gr.Greeting = "Hello!"
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	render.JSON(w, req, &gr)
}

func main() {
	docFlag := flag.Bool("d", false, "print API documentation")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", Index)
	r.Get("/{lang}", Greet)

	if *docFlag {
		fmt.Println(docgen.JSONRoutesDoc(r))
		os.Exit(0)
	}

	log.Println("Listening on http://0.0.0.0:9999/")
	log.Fatal(http.ListenAndServe(":9999", r))
}
