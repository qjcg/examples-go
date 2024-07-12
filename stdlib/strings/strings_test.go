package strings

import (
	"strings"
	"testing"
)

func TestToUpper(t *testing.T) {
	f := func(input, want string) {
		t.Helper()

		got := strings.ToUpper(input)
		if got != want {
			t.Fatalf("want %q got %q", want, got)
		}
	}

	f("", "")
	f("hello", "HELLO")
	f("a multi-word sentence", "A MULTI-WORD SENTENCE")
}
