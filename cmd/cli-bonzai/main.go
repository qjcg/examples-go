package main

import (
	"fmt"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/comp"
	"github.com/rwxrob/bonzai/vars"
)

var cmdRoot = &bonzai.Cmd{
	Name:  "example",
	Alias: "ex|eg",
	Vers:  "0.1.0",
	Usage: "[-h|--help|help]",
	Short: "an example command with Bonzai",
	Long:  "",
	Comp:  comp.Cmds,
	Cmds: []*bonzai.Cmd{
		help.Cmd,
		vars.Cmd,
	},
	Def: help.Cmd,
	Do: func(cmd *bonzai.Cmd, _ ...string) error {
		fmt.Printf("executing command: %s\n", cmd.Name)

		return nil
	},
	Init: func(caller *bonzai.Cmd, _ ...string) error {
		return nil
	},
	MinArgs: 0,
	MaxArgs: 0,
	NumArgs: 0,
	NoArgs:  false,
}

func Main() {
	cmdRoot.Exec()
}

func main() {
	Main()
}
