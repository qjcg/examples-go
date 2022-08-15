package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Subcommand: serve
	cmdServe = &cobra.Command{
		Use:   "serve",
		Short: "Serve our web application",
		Run: func(cmd *cobra.Command, args []string) {
			Serve(port)
		},
	}
)

// Flag vars.
var (
	port int
)

// Add a flag for port, and add our subcommand to the root command.
func init() {
	cmdServe.PersistentFlags().IntVarP(&port, "port", "p", 8080, "port to listen on")

	rootCmd.AddCommand(cmdServe)
}

// Serve will start a server listening on the provided port.
func Serve(port int) {
	fmt.Printf("Listening on port: %d\n", port)
}
