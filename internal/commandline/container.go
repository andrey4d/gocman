package commandline

import (
	"fmt"
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

	fmt.Println(args)

	cAtrs := containers.ContainerAttr{
		Command_name:   args[0],
		Arguments:      args[1:],
		Container_name: uuid.New().String(),
		Root:           "fakeroot",

		OvfsRoot: &containers.OvfsMountCfg{
			Lowerdir: []string{"ovfs/l1", "ovfs/l2"},
			Upperdir: "ovfs/upper",
			Workdir:  "ovfs/work",
			Target:   "fakeroot",
			SELebel:  "",
		},
	}
	fmt.Println(cAtrs.Command_name)
	fmt.Println(cAtrs.Arguments)
	containers.Container(cAtrs)
}
