package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	flagPort := flag.Int("p", 8888, "port to listen on")
	flag.Parse()

	http.HandleFunc("/", handleIndex)
	log.Printf("Listening on port %d\n", *flagPort)
	portString := fmt.Sprintf(":%d", *flagPort)
	log.Fatal(http.ListenAndServe(portString, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
