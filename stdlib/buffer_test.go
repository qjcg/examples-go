package stdlib

import (
	"bytes"
	"testing"
)

func TestBufferWriteString(t *testing.T) {
	f := func(s string) {
		t.Helper()

		var buf bytes.Buffer
		buf.WriteString(s)

		want := s
		got := buf.String()

		if want != got {
			t.Fatalf("want %q got %q", want, got)
		}

		t.Logf("buffer contents: %q", got)
	}

	f("")
	f(" ")
	f("test\n")
	f("foo bar baz")
	f(`a


weird

            multi-line



  ____      string`)
}
