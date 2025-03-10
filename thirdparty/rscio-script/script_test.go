package script_test

import (
	"context"
	"testing"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

func Test_scripttest(t *testing.T) {
	ctx := context.Background()
	engine := script.NewEngine()

	scripttest.Test(
		t, ctx, engine,
		[]string{
			"PATH=/usr/bin",
		},
		"testdata/*.txtar")
}
