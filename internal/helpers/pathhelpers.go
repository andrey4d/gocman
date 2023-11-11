package helpers

import (
	"os"
	"path/filepath"
)

func GetAbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	pwd, err := os.Getwd()

	ErrorHelperPanicWithMessage(err, "GetAbsPath()")
	return filepath.Join(pwd, path)
}
