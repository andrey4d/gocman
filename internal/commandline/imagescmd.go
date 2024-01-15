/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"github.com/andrey4d/gocman/internal/config"
	"github.com/andrey4d/gocman/internal/containers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	config.InitContainersHome(viper.GetString(flag_containers.Name))
	containers.ListImages()
}
