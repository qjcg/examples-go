package main

import (
	"context"
	"testing"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

func Test_migrate(t *testing.T) {
	scripttest.Test(
		t,
		context.Background(),
		script.NewEngine(),
		[]string{"PATH=/usr/bin"},
		"testdata/*.txtar",
	)
}
