package cmd

import (
	"strconv"

	"github.com/TwiN/go-color"
	"github.com/enkodr/memoir/memoir"
	"github.com/spf13/cobra"
)

var undoCommand = &cobra.Command{
	Use:   "undo <task_id>",
	Short: "Mark a task as undone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := memoir.Load()
		if err != nil {
			println(color.InRed(err.Error()))
		}

		i, err := strconv.Atoi(args[0])
		if err != nil {
			println(color.InRed("The id needs to be an integer"))
		}

		if i < 0 || i > len(m.DailyTasks) {
			println(color.InRed("The task doesn't exist"))
		}

		m.DailyTasks[i-1].Done = false

		m.SaveTasks()
	},
}

func init() {
	rootCommand.AddCommand(undoCommand)
}
