// A basic HTTP server example.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

// greetResponse represents the data returned as a JSON response.
type greetResponse struct {
	Greeting string `json:"greeting"`
}

func main() {
	flagPort := flag.Int("p", 8888, "port to listen on")
	flag.Parse()

	http.HandleFunc("/", handleIndex)
	log.Printf("Listening on port %d\n", *flagPort)
	portString := fmt.Sprintf(":%d", *flagPort)
	log.Fatal(http.ListenAndServe(portString, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	name := strings.Title(strings.Trim(html.EscapeString(r.URL.Path), "/"))
	greeting := fmt.Sprintf("Hello, %s!", name)
	if name == "" {
		greeting = "Hello!"
	}

	data := greetResponse{Greeting: greeting}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR\n"))
	}
}
