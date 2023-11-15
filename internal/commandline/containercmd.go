/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"github.com/spf13/cobra"

	"godman/internal/containers"
)

var containerCmd = &cobra.Command{
	Use: "container",
	Run: container,
}

func init() {
	rootCmd.AddCommand(containerCmd)
}

func container(cmd *cobra.Command, args []string) {

	containers.Container(args)
}
