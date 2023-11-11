package commandline

import (
	"godman/internal/helpers"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Container from scratch.",
}

func Execute() {
	helpers.ErrorHelperPanicWithMessage(rootCmd.Execute(), "cobra root menu")
}
