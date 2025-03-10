package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"greet": main,
	})
}

func TestBasicsGreeter(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
