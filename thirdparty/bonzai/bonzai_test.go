package bonzai

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/rwxrob/bonzai"
)

func TestCommand(t *testing.T) {
	var buf bytes.Buffer
	want := "foo stuff"

	testCommand := &bonzai.Cmd{
		Name: "foo",
		Do: func(_ *bonzai.Cmd, _ ...string) error { // note conventional _
			fmt.Fprintf(&buf, "%s", want)
			return nil
		},
	}

	if err := testCommand.Do(nil); err != nil {
		t.Fatalf("failed to call command: %v", err)
	}

	got := buf.String()
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
