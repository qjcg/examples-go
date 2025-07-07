package embed

import (
	"bytes"
	"embed"
	"io/fs"
	"testing"
)

//go:embed testdata/greeting.txt
var greeting string

//go:embed testdata/greeting.txt
var greetingBytes []byte

//go:embed testdata/greetings
var greetingsFS embed.FS

func Test_embedString(t *testing.T) {
	got := greeting
	want := "hello\n"
	if got != want {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func Test_embedByteSlice(t *testing.T) {
	got := greetingBytes
	want := []byte("hello\n")
	if !bytes.Equal(got, want) {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func Test_embedFS(t *testing.T) {
	t.Run("ReadFile", func(t *testing.T) {
		got, err := greetingsFS.ReadFile("testdata/greetings/fr.txt")
		if err != nil {
			t.Fatal(err)
		}

		want := []byte("bonjour\n")
		if !bytes.Equal(got, want) {
			t.Fatalf("want: %s, got: %s", want, got)
		}
	})

	t.Run("SubFS", func(t *testing.T) {
		subFS, err := fs.Sub(greetingsFS, "testdata/greetings")

		got, err := fs.ReadFile(subFS, "fr.txt")
		if err != nil {
			t.Fatal(err)
		}

		want := []byte("bonjour\n")
		if !bytes.Equal(got, want) {
			t.Fatalf("want: %s, got: %s", want, got)
		}
	})
}
