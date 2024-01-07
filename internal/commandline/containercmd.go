/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"godman/internal/config"
	"godman/internal/containers"
)

var containerCmd = &cobra.Command{
	Use:    "container",
	Run:    container,
	Hidden: true,
}

func init() {
	rootCmd.AddCommand(containerCmd)
}

func container(cmd *cobra.Command, args []string) {
	config.InitContainersHome(viper.GetString(flag_containers.Name))
	containers.Container(args)
}
