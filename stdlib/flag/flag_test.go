package flag

import (
	"flag"
	"testing"
)

func TestFlagSet_Parse(t *testing.T) {
	type Conf struct {
		Format string
		Number uint
	}

	var conf Conf
	want := "toml"

	fset := flag.NewFlagSet("test", flag.ExitOnError)
	fset.StringVar(&conf.Format, "f", "yaml", "config file format")
	fset.UintVar(&conf.Number, "n", 42, "number of items")

	// Parse arguments that override the default.
	err := fset.Parse([]string{
		"-f", want,
	})
	if err != nil {
		t.Fatalf("error parsing flags: %v", err)
	}

	got := conf.Format
	if got != want {
		t.Fatalf("want %v got %v", want, got)
	}
}
