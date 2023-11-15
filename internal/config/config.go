/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package config

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Config struct {
		ContainersPath string
		ContainersTemp string
		ImageDbPath    string
		Permissions    fs.FileMode
	}
)

type Container struct {
	BasePath      string      `yaml:"base_path"`
	TempPath      string      `yaml:"temp_path"`
	ContainerPath string      `yaml:"container_path"`
	ContainerPerm fs.FileMode `yaml:"container_perm"`
}

type ConfigOld struct {
	Container Container `yaml:"container"`
}

func InitConfig(configPath string) *ConfigOld {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Printf("Config file %s doesn't exist.", configPath)

	}

	file, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("can't read config.")
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var cfg ConfigOld

	decoder.Decode(&cfg)
	fmt.Printf("%v\n", cfg.Container.BasePath)
	return &cfg
}
