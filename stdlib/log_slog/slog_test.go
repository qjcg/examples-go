package slog

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	want := "hello"

	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{}))
	logger.Info(want)

	got := buf.String()
	if !strings.Contains(got, want) {
		t.Fatalf("got %q does not contain want %q", want, got)
	}
}
