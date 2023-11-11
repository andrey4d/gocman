package commandline

import (
	"godman/internal/containers"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use: "container",
	Run: container,
}

func init() {
	rootCmd.AddCommand(containerCmd)
}

func container(cmd *cobra.Command, args []string) {

	var ovfsRoot = &containers.OvfsMountCfg{
		Lowerdir: []string{"ovfs/l1", "ovfs/l2"},
		Upperdir: "ovfs/upper",
		Workdir:  "ovfs/work",
		Target:   "fakeroot",
		SELebel:  "",
	}

	cAtrs := containers.ContainerAttr{
		Command_name:   args[0],
		Arguments:      args[1:],
		Container_name: uuid.New().String(),
		OvfsRoot:       ovfsRoot,
	}

	containers.Container(cAtrs)
}
