/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/andrey4d/gocman/cmd/cmount/cmountconfig"
	"github.com/andrey4d/gocman/internal/config"
	"github.com/andrey4d/gocman/internal/containers"
	"github.com/andrey4d/gocman/internal/helpers"
)

func main() {
	image_name := os.Args[1]
	mount_point := os.Args[2]

	cmountconfig.InitContainerConfig("containers")
	config.CheckImagesPath()

	imageId := containers.DownloadImage(image_name)

	overlayMountCfg := containers.OvfsMountCfg{
		Lowerdir:    nil,                                                   // overlay/l/<ZLR4NWYDXWB5LCOBDH7WAGVYDI> --> storage/overlay/<id>/diff
		Upperdir:    config.Config.GetContainersPath() + "/ovfs/diff",      // storage/overlay/<id>/diff
		Workdir:     config.Config.GetContainersPath() + "/ovfs/work",      // storage/overlay/<id>/work
		Target:      config.Config.GetContainersPath() + "/" + mount_point, // storage/overlay/<id>/merged
		Permeations: config.Config.GetPermissions(),
		SELebel:     "",
	}
	overlayMountCfg.CreateDirectoryStructure()
	helpers.CheckError(containers.MountOvfs(imageId, &overlayMountCfg), "containe() mount overlay")
	log.Printf("mounted to %s\n", helpers.GetAbsPath(overlayMountCfg.Target))
}
