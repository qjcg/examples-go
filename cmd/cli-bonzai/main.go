package main

import (
	"fmt"
	"os"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var cmdRoot = &bonzai.Cmd{
	Name:        "example",
	Short:     "An example command with Bonzai",
	Usage:       "[-h|--help|help]",
	Long: "",
	Cmds: []*bonzai.Cmd{
		help.Cmd,
	},
	Comp:    nil,
	Do: func(cmd *bonzai.Cmd, _ ...string) error {
		fmt.Printf("executing command: %s\n", cmd.Name)

		return nil
	},
	Init: func(caller *bonzai.Cmd, _ ...string) error {
		return nil
	},
	MinArgs:  0,
	MaxArgs:  0,
	NumArgs:  0,
	NoArgs:   false,
}

func Main() int {
	cmdRoot.Run()

	return 0
}

func main() {
	os.Exit(Main())
}
