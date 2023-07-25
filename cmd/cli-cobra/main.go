package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobrademo",
	Short: "A demo application using cobra",
}

func Main() int {
	rootCmd.Execute()
	return 0
}

func main() {
	os.Exit(Main())
}
