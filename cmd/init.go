package cmd

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/enkodr/memoir/memoir"
	"github.com/spf13/cobra"
)

var (
	local bool
)

// initCommand is a command line action to initialise memoir
var initCommand = &cobra.Command{
	Use:                   "init [path]",
	Short:                 "Initilises memoir in the current directory",
	Args:                  cobra.MaximumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		// Initialise the structure
		m, err := memoir.Init()
		if err != nil {
			println(color.InRed(err.Error()))
			return
		}
		println(color.InGreen(fmt.Sprintf("memoir is succesfully inicialised in %q", m.Path)))
	},
}

func init() {
	rootCommand.AddCommand(initCommand)
}
