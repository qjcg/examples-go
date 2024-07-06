package main_test

import (
	"bytes"
	"strings"
	"testing"

	main "github.com/qjcg/go-examples/cmd/basics-concurrency-visualization"
)

func TestUp(t *testing.T) {
	f := func(want string) {
		t.Helper()
		var b bytes.Buffer
		main.Up(&b)
		got := b.String()

		if !strings.Contains(got, want) {
			t.Fatal("desired substring not in output")
		}
	}

	f("")
	f(" ")
	f("U42")
}
