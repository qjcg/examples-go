package stdlib

import (
	"flag"
	"testing"
)

type Conf struct {
	Format string
	Number uint
}

func TestFlagSet(t *testing.T) {
	var conf Conf
	want := "toml"

	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.StringVar(&conf.Format, "f", "yaml", "config file format")
	fs.UintVar(&conf.Number, "n", 42, "number of items")
	fs.Parse([]string{
		"-f", want,
	})

	got := conf.Format
	if got != want {
		t.Fatalf("want %v got %v", want, got)
	}

	t.Logf("%v", conf)
}
