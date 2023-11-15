/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package commandline

import (
	"fmt"
	"godman/internal/config"
	"godman/internal/helpers"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config/config.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home + "/config")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		initContainerConfig()
		helpers.CheckImagesPath()

	}
}

func initContainerConfig() {
	config.Config.SetContainersPath(helpers.GetAbsPath(viper.GetString("container.container_path")))
	config.Config.SetContainersTemp(fmt.Sprintf("%s/%s", config.Config.GetContainersPath(), viper.GetString("container.temp_path")))
	config.Config.SetOverlayImage(fmt.Sprintf("%s/storage/overlay-images", config.Config.GetContainersPath()))
	config.Config.SetImageDbPath(fmt.Sprintf("%s/images.json", config.Config.GetOverlayImage()))
	config.Config.SetPermissions(fs.FileMode(viper.GetUint32("container.container_perm")))
	config.Config.SetOverlayLinkDir(fmt.Sprintf("%s/storage/overlay/l", config.Config.GetContainersPath()))
}
