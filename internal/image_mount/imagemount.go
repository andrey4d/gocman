/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package imagemount

import (
	"github.com/andrey4d/gocman/internal/config"
	"github.com/andrey4d/gocman/internal/containers"
	"github.com/andrey4d/gocman/internal/helpers"

	"github.com/sirupsen/logrus"
)

func ImageMountToDir(args []string) {

	image_name := args[0]
	mount_point := args[1]

	logrus.Infof("mount image %s to %s\n", image_name, mount_point)

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
	logrus.Printf("mounted to %s\n", helpers.GetAbsPath(overlayMountCfg.Target))
}
