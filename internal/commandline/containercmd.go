package commandline

import (
	"fmt"

	"github.com/spf13/cobra"

	"godman/internal/config"
	"godman/internal/containers"
	"godman/internal/helpers"
)

var containerCmd = &cobra.Command{
	Use: "container",
	Run: container,
}

func init() {
	rootCmd.AddCommand(containerCmd)
}

func container(cmd *cobra.Command, args []string) {
	cfg := config.InitConfig("config/config.yaml")
	cId := helpers.CreateContainerID(16)

	containerBaseDir := fmt.Sprintf("%s/storage/overlay/%s", containers.GetAbsPath(cfg.Container.ContainerPath), cId)

	var ovfsRoot = &containers.OvfsMountCfg{
		Lowerdir:    []string{fmt.Sprintf("%s/l", containerBaseDir)}, // overlay/l/<ZLR4NWYDXWB5LCOBDH7WAGVYDI> --> storage/overlay/<id>/diff
		Upperdir:    fmt.Sprintf("%s/diff", containerBaseDir),        // storage/overlay/<id>/diff
		Workdir:     fmt.Sprintf("%s/work", containerBaseDir),        // storage/overlay/<id>/work
		Target:      fmt.Sprintf("%s/merged", containerBaseDir),      // storage/overlay/<id>/merged
		Permeations: cfg.Container.ContainerPerm,
		SELebel:     "",
	}

	cAtrs := containers.ContainerAttr{
		Command_name:   args[0],
		Arguments:      args[1:],
		Container_name: cId,
		OvfsRoot:       ovfsRoot,
	}

	containers.Container(cAtrs)
}
