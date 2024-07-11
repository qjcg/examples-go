package stdlib

import (
	"bytes"
	"log/slog"
	"regexp"
	"testing"
)

func TestSlog(t *testing.T) {
	var buf bytes.Buffer
	want := "hello"
	REwant := regexp.MustCompile(want)

	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{}))
	logger.Info(want)

	got := buf.String()
	if !REwant.MatchString(got) {
		t.Fatalf("wanted match for %v got %v", REwant.String(), got)
	}
	t.Logf("\nwant: %v\ngot: %v", want, got)
}
