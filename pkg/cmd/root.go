package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bri",
	Short: "A simple CLI for redundant tasks.",
	Long:  "TODO",
}

// ----------------------------------------------------
// Command Registry
// register any new commands or command groups here
var commandRegistry = []*cobra.Command{
	wordleCmd(),
	timeDialtionCmd(),
}

// ----------------------------------------------------
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	for i := 0; i < len(commandRegistry); i++ {
		rootCmd.AddCommand(commandRegistry[i])
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
