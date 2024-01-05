/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"godman/internal/helpers"
	imagemount "godman/internal/image_mount"

	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Short: "mount <image> <mount dir>",
	Use:   "mount",
	Run:   mount,
}

func init() {
	rootCmd.AddCommand(mountCmd)
}

func mount(_ *cobra.Command, args []string) {
	if len(args) == 0 || len(args[0]) == 0 {
		helpers.FatalHelperLog("command not specified")

	}

	imagemount.ImageMountToDir(args)
}
