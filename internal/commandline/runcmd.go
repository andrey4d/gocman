package commandline

import (
	"godman/internal/helpers"
	"godman/internal/starter"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Short: "run <cmd> in container",
	Use:   "run",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(_ *cobra.Command, args []string) {
	if len(args) == 0 || len(args[0]) == 0 {
		helpers.FatalHelperLog("command not specified")

	}

	starter.Run(args)
}
