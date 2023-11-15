/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package helpers

import (
	"fmt"
	"godman/internal/config"
)

func CheckImagesPath() {
	fmt.Println("Init images paths..")
	perm := config.Config.GetPermissions()
	MakeDirAllIfNotExists(config.Config.GetContainersTemp(), perm)
	MakeDirAllIfNotExists(config.Config.GetOverlayLinkDir(), perm)
	MakeDirAllIfNotExists(config.Config.GetOverlayImage(), perm)
}
