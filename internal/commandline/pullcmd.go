/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"github.com/andrey4d/gocman/internal/config"
	"github.com/andrey4d/gocman/internal/containers"
	"github.com/andrey4d/gocman/internal/helpers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	config.InitContainersHome(viper.GetString(flag_containers.Name))
	containers.DownloadImage(args[0])
}
