/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cmountconfig

import (
	"godman/internal/config"
	"godman/internal/helpers"
	"io/fs"
)

func InitContainerConfig(mount_path string) {

	config.Config.SetContainersPath(helpers.GetAbsPath(mount_path))

	config.Config.SetContainersTemp(config.Config.GetContainersPath() + "/tmp")
	config.Config.SetOverlayImage(config.Config.GetContainersPath() + "/storage/overlay-images")
	config.Config.SetImageDbPath(config.Config.GetOverlayImage() + "/images.json")
	config.Config.SetPermissions(fs.FileMode(0755))
	config.Config.SetOverlayLinkDir(config.Config.GetContainersPath() + "/storage/overlay/l")
	config.Config.SetOverlayDir(config.Config.GetContainersPath() + "/storage/overlay")
}
