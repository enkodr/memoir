package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "memoir",
	Short: "A client tool to manage daily tasks",
}

func Execute() {
	rootCommand.Execute()
}
