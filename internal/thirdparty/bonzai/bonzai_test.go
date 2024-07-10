package bonzai

import (
	"bytes"
	"fmt"
	"testing"

	Z "github.com/rwxrob/bonzai/z"
)

func TestCommand(t *testing.T) {
	var buf bytes.Buffer
	want := "foo stuff"

	testCommand := &Z.Cmd{
		Name: "foo",
		Call: func(_ *Z.Cmd, _ ...string) error { // note conventional _
			fmt.Fprintf(&buf, want)
			return nil
		},
	}

	if err := testCommand.Call(nil); err != nil {
		t.Fatalf("failed to call command: %v", err)
	}

	got := buf.String()
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
