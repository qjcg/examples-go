package slog

import (
	"bytes"
	"regexp"
	"testing"

	"golang.org/x/exp/slog"
)

func TestSlog(t *testing.T) {
	var buf bytes.Buffer
	want := "hello"
	REwant := regexp.MustCompile(want)

	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{}))
	logger.Info(want)

	got := buf.String()
	t.Logf("\nwant: %v\ngot: %v", want, got)
	if !REwant.MatchString(got) {
		t.Fatalf("wanted match for %v got %v", REwant.String(), got)
	}
}
