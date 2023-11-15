/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package starter

import (
	"fmt"
	"godman/internal/helpers"
	"os"
	"os/exec"
	"syscall"
)

func Run(args []string) {
	fmt.Println("Hello godman!")

	command := append([]string{"container"}, args...)

	cmd := exec.Command("/proc/self/exe", command...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	cmd.SysProcAttr.Cloneflags = syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER

	cmd.SysProcAttr.UidMappings = []syscall.SysProcIDMap{
		{
			ContainerID: 0,
			HostID:      1000,
			Size:        1,
		},
	}

	cmd.SysProcAttr.GidMappings = []syscall.SysProcIDMap{
		{
			ContainerID: 0,
			HostID:      1000,
			Size:        1,
		},
	}

	fmt.Printf("Starter PID: %d\n", os.Getpid())
	helpers.ErrorHelperPanicWithMessage(cmd.Run(), "Run() ")
}
