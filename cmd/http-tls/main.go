package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	cert := flag.String("c", "", "TLS Cert file")
	key := flag.String("k", "", "TLS Key file")
	flag.Parse()

	if *cert == "" || *key == "" {
		log.Fatal("Invalid TLS cert and/or key provided!")
	}

	http.HandleFunc("/", Handler)
	log.Println("Listening on https://0.0.0.0:9999/")
	log.Fatal(http.ListenAndServeTLS(":9999", *cert, *key, nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing request form!", http.StatusInternalServerError)
		return
	}

	// Timeout defaults to 1 second. A timeout value can also be provided
	// via POST request in Milliseconds.
	timeout := time.Second
	if r.Method == "POST" {
		if v := r.Form.Get("timeout"); v != "" {
			t, err := strconv.Atoi(v)
			if err != nil {
				http.Error(w, "Error parsing POSTED timeout value!", http.StatusBadRequest)
				return
			}
			timeout = time.Duration(t) * time.Millisecond
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	var resp Response
	for {
		select {
		case <-ctx.Done():
			return

		default:
			resp.Time = time.Now()
			if err := enc.Encode(&resp); err != nil {
				http.Error(w, "Error encoding data!", http.StatusInternalServerError)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}
}

type Response struct {
	Time time.Time
}
