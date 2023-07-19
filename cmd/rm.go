package cmd

import (
	"strconv"

	"github.com/TwiN/go-color"
	"github.com/enkodr/memoir/memoir"
	"github.com/spf13/cobra"
)

var rmCommand = &cobra.Command{
	Use:     "rm \"<task_id>\"",
	Short:   "Remove a daily task",
	Aliases: []string{"del", "delete"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := memoir.Load()
		if err != nil {
			println(color.InRed(err.Error()))
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			println(color.InRed("The id needs to be an integer"))
		}

		err = m.DeleteTask(id)

		if err != nil {
			println(color.InRed("error deleting task"))
		}
	},
}

func init() {
	rootCommand.AddCommand(rmCommand)
}
