package cmd

import (
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/TwiN/go-color"
	"github.com/enkodr/memoir/memoir"
	"github.com/spf13/cobra"
)

var editCommand = &cobra.Command{
	Use:                "edit [nubmer_of_days]",
	Short:              "Edit the tasks file in default editor",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		dateArg := "0"
		if len(args) > 0 {
			dateArg = args[0]
		}

		days, err := strconv.Atoi(dateArg)
		if err != nil {
			println(color.InRed("The argument must be the number of days"))
		}

		date := time.Now().Add(time.Duration(days) * time.Hour * 24)
		m, err := memoir.LoadFromDate(date)
		if err != nil {
			println(color.InRed(err.Error()))
		}

		editor := os.Getenv("EDITOR")
		c := exec.Command(editor, m.TasksFile)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err = c.Run()
		if err != nil {
			println(color.InRed(err.Error()))
		}

	},
}

func init() {
	rootCommand.AddCommand(editCommand)
}
