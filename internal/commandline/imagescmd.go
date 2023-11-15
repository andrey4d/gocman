/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"godman/internal/containers"

	"github.com/spf13/cobra"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Show image",
	Run:   images,
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}

func images(_ *cobra.Command, _ []string) {

  containers .ListImages()
}
