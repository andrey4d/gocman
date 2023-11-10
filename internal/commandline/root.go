package commandline

import (
	"godman/internal/handlers"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Container from scratch.",
}

func Execute() {
	handlers.ErrorHandlerPanicWithMessage(rootCmd.Execute(), "cobra root menu")
}
