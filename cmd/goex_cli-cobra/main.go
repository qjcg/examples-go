package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "cobrademo",
	Short: "A demo application using cobra",
}

func main() {
	rootCmd.Execute()
}
