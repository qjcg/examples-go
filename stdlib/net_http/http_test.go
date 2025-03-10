package net_http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodPatterns_GET(t *testing.T) {
	wantedGetOutput := "Successful GET!"

	mux := http.NewServeMux()

	indexGetHandler := func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, wantedGetOutput)
		if err != nil {
			t.Fatalf("error writing index content: %v", err)
		}
	}

	mux.HandleFunc("GET /", indexGetHandler)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	resp := w.Result()
	_, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	if w.Body.String() != wantedGetOutput {
		t.Errorf("Expected response body '%s' but got '%s'", wantedGetOutput, w.Body.String())
	}
}
