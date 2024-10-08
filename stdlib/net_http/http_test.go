package net_http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodPatterns(t *testing.T) {
	mux := http.NewServeMux()

	indexHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "all good!")
	}

	mux.Handle("GET /", indexHandler)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	if w.Body.String() != "all good!" {
		t.Errorf("Expected response body 'all good!' but got '%s'", w.Body.String())
	}
}
