package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TwiN/go-color"
	"github.com/enkodr/memoir/memoir"
	"github.com/spf13/cobra"
)

var showCommand = &cobra.Command{
	Use:                "show [number_of_days]",
	Short:              "Show a list of the daily task",
	DisableFlagParsing: true,
	ArgAliases:         []string{"nb"},
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

		dateStr := fmt.Sprintf("Date: %v", m.Today)
		println(color.InBlue(dateStr))

		for i, task := range m.DailyTasks {
			print(color.InYellow(i+1), " - ")
			if task.Done {
				print(color.InWhite("["), color.InGreen("X"), color.InWhite("]"), " ")
			} else {
				print(color.InWhite("[ ]"), " ")
			}
			println(task.Title)
		}
	},
}

func init() {
	rootCommand.AddCommand(showCommand)
}
