package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

// FIXME
func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"example": Main,
	})
}

func TestBasicsGreeter(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
