package cmd

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
	"github.com/enkodr/memoir/memoir"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:                "list",
	Short:              "Show a list of all days with tasks",
	DisableFlagParsing: true,
	ArgAliases:         []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {

		m, err := memoir.Load()
		if err != nil {
			println(color.InRed(err.Error()))
		}

		today := time.Now()
		for _, d := range m.GetAllDailies() {
			date, err := time.Parse("20060102", d)
			if err == nil {
				now := date.Format("2006/01/02")
				print(color.InBlue(now))
				diff := date.Sub(today)
				days := int(diff.Hours() / 24)
				println(color.InYellow(fmt.Sprintf(" (%d)", days)))
			}
		}

	},
}

func init() {
	rootCommand.AddCommand(listCommand)
}
