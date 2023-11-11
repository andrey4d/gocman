package commandline

import (
	"github.com/spf13/cobra"

	"godman/internal/helpers"
)

var rootCmd = &cobra.Command{
	Short: "Container from scratch.",
}

func Execute() {
	helpers.ErrorHelperPanicWithMessage(rootCmd.Execute(), "root menu")

}
