/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"godman/internal/containers"
	"godman/internal/helpers"

	"github.com/spf13/cobra"
)

// pullCmd represents the images command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull image",
	Run:   pull,
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func pull(_ *cobra.Command, args []string) {
	if len(args) == 0 || len(args[0]) == 0 {
		helpers.FatalHelperLog("image name not specified")

	}

	containers.DownloadImage(args[0])
}
