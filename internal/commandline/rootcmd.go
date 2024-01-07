/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type flag struct {
	Name         string
	Short        string
	DefaultValue string
	Descriptions string
}

var (
	cfgFile         string
	flag_containers = flag{
		Name:         "container",
		Short:        "c",
		DefaultValue: "./containers",
		Descriptions: "containers path or set env CONTAINERS",
	}

	rootCmd = &cobra.Command{
		Short: "Container from scratch.",
		Run:   runRoot,
	}
)

func runRoot(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(cobraInit)
	// add config flag
	rootCmd.PersistentFlags().StringVarP(&cfgFile, flag_containers.Name, flag_containers.Short, flag_containers.DefaultValue, flag_containers.Descriptions)
	viper.BindPFlag(flag_containers.Name, rootCmd.PersistentFlags().Lookup(flag_containers.Name))
}

func cobraInit() {
	viper.AutomaticEnv()
}
