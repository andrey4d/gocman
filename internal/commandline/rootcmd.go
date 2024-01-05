/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"godman/internal/config"

	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Short: "Container from scratch.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	container_path := "containers"
	config.InitContainerConfig(container_path)
	config.CheckImagesPath()
}
