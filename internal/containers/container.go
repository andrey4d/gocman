/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package containers

import (
	"fmt"
	"godman/internal/config"
	"godman/internal/helpers"
	"os"
	"os/exec"

	"syscall"
)

func Container(args []string) {

	cId := helpers.CreateContainerID(16)

	image_name := args[0]
	imageId := DownloadImage(image_name)

	// imageId, err := GetIdByName(image_name)

	// helpers.CheckError(err, "container() get image id")

	command_name := args[1]
	arguments := args[2:]

	var overlayMountCfg = OvfsMountCfg{
		Lowerdir:    nil,                                                                                    // overlay/l/<ZLR4NWYDXWB5LCOBDH7WAGVYDI> --> storage/overlay/<id>/diff
		Upperdir:    fmt.Sprintf("%s/storage/containers/%s/diff", config.Config.GetContainersPath(), cId),   // storage/overlay/<id>/diff
		Workdir:     fmt.Sprintf("%s/storage/containers/%s/work", config.Config.GetContainersPath(), cId),   // storage/overlay/<id>/work
		Target:      fmt.Sprintf("%s/storage/containers/%s/merged", config.Config.GetContainersPath(), cId), // storage/overlay/<id>/merged
		Permeations: config.Config.GetPermissions(),
		SELebel:     "",
	}
	overlayMountCfg.CreateDirectoryStructure()

	fmt.Printf("change root to %s\n", helpers.GetAbsPath(overlayMountCfg.Target))

	helpers.CheckError(MountOvfs(imageId, &overlayMountCfg), "containe() mount overlay")

	// MountProc(overlayMountCfg.Target)

	// helpers.CheckError(MountRoot(overlayMountCfg.Target), "container() pivot root")

	// helpers.CheckError(syscall.Chdir("/"), "change dir")

	cmd := exec.Command(command_name, arguments...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PATH=/bin:/sbin:/usr/bin:/usr/sbin"}

	setHostname(cId)

	fmt.Printf("Container PID: %d\n", os.Getpid())
	helpers.CheckError(cmd.Run(), "run container")

	helpers.CheckError(UmountProc(), "umount /proc")
}

func setHostname(hostname string) {
	helpers.CheckError(syscall.Sethostname([]byte(hostname)), "set container name")
}
