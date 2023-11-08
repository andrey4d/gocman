package main

import (
	"fmt"
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
			Container_name: "container",
			Root:           "fakeroot",
		}

		containers.Child(cAtrs)
	default:
		panic("help: use run <cmd>")
	}
}
