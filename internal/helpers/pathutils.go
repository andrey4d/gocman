/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	pwd, err := os.Getwd()
	ErrorHelperPanicWithMessage(err, "get work dir")

	return filepath.Join(pwd, path)
}

func GetTempPath(tmpPath string) string {

	path, err := os.MkdirTemp(tmpPath, "*")
	ErrorHelperPanicWithMessage(err, "make temp dir")
	return path
}

func MakeTempPath(path string, id string) string {
	tmp := fmt.Sprintf("%s/%s", GetAbsPath(path), id)
	err := os.MkdirAll(tmp, 0755)
	ErrorHelperPanicWithMessage(err, "can't make temp dir")
	return tmp
}
