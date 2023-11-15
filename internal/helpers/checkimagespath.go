/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package helpers

import (
	"fmt"

	"io/fs"

	"github.com/spf13/viper"
)

func CheckImagesPath() {
	fmt.Println("Init images paths..")
	perm := fs.FileMode(viper.GetUint32("container.container_perm"))
	containersPath := GetAbsPath(viper.GetString("container.container_path"))
	MakeDirAllIfNotExists(fmt.Sprintf("%s/%s", containersPath, viper.GetString("container.temp_path")), perm)
	MakeDirAllIfNotExists(fmt.Sprintf("%s/storage/overlay/l", containersPath), perm)
	MakeDirAllIfNotExists(fmt.Sprintf("%s/storage/overlay-images", containersPath), perm)
}
