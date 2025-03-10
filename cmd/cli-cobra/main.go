package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobrademo",
	Short: "A demo application using cobra",
}

func Main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Main()
}
