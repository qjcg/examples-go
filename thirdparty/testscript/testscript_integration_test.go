//go:build integration

package testing_testscript

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestRunTestscripts(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "integration",
	})
}
