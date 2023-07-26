package main

import (
	"fmt"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var cmdRoot = &Z.Cmd{
	Name:        "example",
	Summary:     "An example command with Bonzai",
	Usage:       "[-h|--help|help]",
	License:     "GPL-3.0",
	Description: "",
	Dynamic:     map[string]any{},
	Source:      "https://github.com/qjcg/go-examples",
	Issues:      "https://github.com/qjcg/go-examples/issues",
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Params:  []string{},
	Hidden:  []string{},
	VarDefs: map[string]string{},
	Comp:    nil,
	Call: func(cmd *Z.Cmd, _ ...string) error {
		fmt.Printf("executing command: %s\n", cmd.Name)

		return nil
	},
	Init: func(caller *Z.Cmd, _ ...string) error {
		return nil
	},
	Input:    nil,
	MinArgs:  0,
	MaxArgs:  0,
	NumArgs:  0,
	NoArgs:   false,
	MinParm:  0,
	MaxParm:  0,
	UseConf:  false,
	UseVars:  false,
	ConfVars: false,
}

func Main() int {
	cmdRoot.Run()

	return 0
}

func main() {
	os.Exit(Main())
}
