/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/andrey4d/gocman/cmd/test/layers/cmountconfig"
	"github.com/andrey4d/gocman/internal/config"
	"github.com/andrey4d/gocman/internal/containers"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {
	cmountconfig.InitContainerConfig("containers")
	config.CheckImagesPath()

	id := "bc4ac1228b3cc9e7a67974810d92e6f752c2e8ef9a11efd4007e608579591693"
	// getLinkToLayers(id)
	layerLinks, err := (containers.GetLowerLayersLink(id))
	if err != nil {
		logrus.Fatal(err)
	}
	for _, ll := range layerLinks {
		llf := config.Config.GetOverlayLinkDir() + "/" + ll
		fmt.Println(llf)
	}

}

// // isParent returns if the passed in parent is the direct parent of the passed in layer
// func isParent(id, parent string) bool {

// 	// dir_1 := "containers/storage/overlay"

// 	lowers, err := d.getLowerDirs(id)
// 	if err != nil {
// 		return false
// 	}
// 	if parent == "" && len(lowers) > 0 {
// 		return false
// 	}

// 	parentDir := d.dir(parent)
// 	var ld string
// 	if len(lowers) > 0 {
// 		ld = filepath.Dir(lowers[0])
// 	}
// 	if ld == "" && parent == "" {
// 		return true
// 	}
// 	return ld == parentDir
// }

func getLinkToLayers(id string) []string {

	path := "containers/storage/overlay/" + id
	lowerFile := "lower"
	lowers, err := os.ReadFile(filepath.Join(path, lowerFile))
	if err != nil {
		logrus.Fatal(err)
	}

	out := strings.Split(string(lowers), ":")

	spew.Dump(out)
	return out
}
