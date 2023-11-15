/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package helpers

import (
	"io/fs"
	"os"
)

func MakeDirAllIfNotExists(name string, chmod fs.FileMode) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.MkdirAll(name, chmod); err != nil {
			return err
		}
	}
	return nil
}

func MakeDirIfNotExists(name string, chmod fs.FileMode) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, chmod); err != nil {
			return err
		}
	}
	return nil
}
