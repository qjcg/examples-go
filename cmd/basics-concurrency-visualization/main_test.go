package main_test

import (
	"bytes"
	"strings"
	"testing"

	main "github.com/qjcg/go-examples/cmd/basics-concurrency-visualization"
)

func TestUp(t *testing.T) {
	testCases := []struct {
		description string
		want        string
	}{
		{"basic", "U42"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			var b bytes.Buffer
			main.Up(&b)
			got := b.String()

			if !strings.Contains(got, tc.want) {
				t.Fatal("desired substring not in output")
			}
		})

	}
}
