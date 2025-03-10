package main

import (
	"bytes"
	_ "embed"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocolly/colly/v2"
)

//go:embed testdata/index.html
var testHTMLPage []byte

func newUnstartedTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(testHTMLPage)
		if err != nil {
			log.Fatalf("error writing testHTMLPage: %v", err)
		}
	})

	return httptest.NewUnstartedServer(mux)
}

func newTestServer() *httptest.Server {
	ts := newUnstartedTestServer()
	ts.Start()
	return ts
}

func TestPrintThumbnailURLs(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	c := colly.NewCollector()

	var buf bytes.Buffer
	c = PrintThumbnailURLs(c, &buf)

	if err := c.Visit(ts.URL); err != nil {
		t.Fatal(err)
	}

	out := buf.Bytes()
	if !REImgurURL.Match(out) {
		t.Fatalf("No imgur URL match in %#v", out)
	}

	t.Log(buf.String())
}
