package main

import (
	"context"
	"io"
	"os"
	"testing"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

func Test_scripttest(t *testing.T) {
	ctx := context.Background()
	engine := script.NewEngine()

	engine.Cmds["main"] = script.Command(
		script.CmdUsage{Summary: "Run main()"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			rOut, wOut, err := os.Pipe()
			if err != nil {
				return nil, err
			}
			rErr, wErr, err := os.Pipe()
			if err != nil {
				return nil, err
			}

			os.Stdout = wOut
			os.Stderr = wErr

			main() // this gets captured

			wOut.Close()
			wErr.Close()
			cmdOut, err := io.ReadAll(rOut)
			if err != nil {
				return nil, err
			}
			cmdErr, err := io.ReadAll(rErr)
			if err != nil {
				return nil, err
			}

			return func(s *script.State) (stdout, stderr string, err error) {
				return string(cmdOut), string(cmdErr), nil
			}, nil
		},
	)

	scripttest.Test(
		t, ctx, engine,
		[]string{
			"PATH=/usr/bin",
		},
		"testdata/*.txtar")
}
