package main

import (
	"fmt"
	"github.com/google/uuid"
	"godman/internal/containers"
	"godman/internal/starter"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s run <cmd>\n", os.Args[0])
		os.Exit(0)
	}

	switch os.Args[1] {
	case "run":
		starter.Run(os.Args)
	case "child":

		cAtrs := containers.ContainerAttr{
			Command_name:   os.Args[2],
			Arguments:      os.Args[3:],
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

		containers.Container(cAtrs)
	default:
		panic("help: use run <cmd>")
	}
}
