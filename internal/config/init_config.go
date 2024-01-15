/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package config

import (
	"io/fs"

	"github.com/andrey4d/gocman/internal/helpers"
	"github.com/sirupsen/logrus"
)

func InitContainerConfig(mount_path string) {

	Config.SetContainersPath(helpers.GetAbsPath(mount_path))
	Config.SetContainersTemp(Config.GetContainersPath() + "/tmp")
	Config.SetOverlayImage(Config.GetContainersPath() + "/storage/overlay-images")
	Config.SetImageDbPath(Config.GetOverlayImage() + "/images.json")
	Config.SetPermissions(fs.FileMode(0755))
	Config.SetOverlayLinkDir(Config.GetContainersPath() + "/storage/overlay/l")
	Config.SetOverlayDir(Config.GetContainersPath() + "/storage/overlay")
}

func CheckImagesPath() {
	logrus.Println("Init images paths..")
	perm := Config.GetPermissions()
	helpers.MakeDirAllIfNotExists(Config.GetContainersTemp(), perm)
	helpers.MakeDirAllIfNotExists(Config.GetOverlayLinkDir(), perm)
	helpers.MakeDirAllIfNotExists(Config.GetOverlayImage(), perm)
}

func InitContainersHome(mount_path string) {
	InitContainerConfig(mount_path)
	CheckImagesPath()
}
