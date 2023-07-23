package cobra

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestCommand(t *testing.T) {
	var buf bytes.Buffer

	testCmd := &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(&buf, "hugo hugo hugo")
		},
	}

	if err := testCmd.Execute(); err != nil {
		t.Fatalf("failed to execute command: %v", err)
	}

	want := "hugo hugo hugo"
	got := buf.String()
	if got != want {
		t.Fatalf("want %v got %v", want, got)
	}
}
