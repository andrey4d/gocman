package config

import (
	"fmt"
	"godman/internal/helpers"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

type container struct {
	BasePath      string      `yaml:"base_path"`
	TempPath      string      `yaml:"temp_path"`
	ContainerPath string      `yaml:"container_path"`
	ContainerPerm fs.FileMode `yaml:"container_perm"`
}

type Config struct {
	Container container `yaml:"container"`
}

func InitConfig(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		helpers.ErrorHelperPanicWithMessage(err, fmt.Sprintf("Config file %s doesn't exist.", configPath))

	}

	file, err := os.Open(configPath)
	if err != nil {
		helpers.ErrorHelperPanicWithMessage(err, "can't read config.")
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var cfg Config

	helpers.ErrorHelperPanicWithMessage(decoder.Decode(&cfg), "")
	fmt.Printf("%v\n", cfg.Container.BasePath)
	return &cfg
}
