package cmd

import (
	"github.com/TwiN/go-color"
	"github.com/enkodr/memoir/memoir"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add \"<task description>\"",
	Short: "Add a daily task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := memoir.Load()
		if err != nil {
			println(color.InRed(err.Error()))
		}

		title := args[0]

		err = m.AddTask(title)

		if err != nil {
			println(color.InRed("error adding task"))
		}
	},
}

func init() {
	rootCommand.AddCommand(addCommand)
}
