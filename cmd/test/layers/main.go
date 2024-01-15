/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {

	id := "66ab603055b6575ebd0d9caa25f073464f2b57dfc7deba3f3f30691bd135bb94"
	getLinkToLayers(id)

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
