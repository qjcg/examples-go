package main

import (
	"embed"
	"strings"
	"testing"
)

//go:embed testdata/greeting.txt
var greeting string

//go:embed testdata/sql
var content embed.FS

func Test_embedString(t *testing.T) {
	got := strings.TrimSpace(greeting)
	want := "hello"
	if got != want {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func Test_embedFS(t *testing.T) {
	got, err := content.ReadFile("testdata/sql/demo.sql")
	if err != nil {
		t.Fatal(err)
	}

	want := "DROP TABLE IF EXISTS users;"
	if !strings.Contains(string(got), want) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
